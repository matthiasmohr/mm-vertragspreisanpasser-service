package priceChangeRuleCollection

import (
	"context"
	"fmt"
	logger "github.com/enercity/lib-logger/v3"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/domain"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/dto"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/repository"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/usecase"
)

type Executer struct {
	store repository.Store
}

func NewExecuter(
	store repository.Store,
) *Executer {
	return &Executer{
		store: store,
	}
}

func (f *Executer) Execute(ctx context.Context, logEntry logger.Entry, req *dto.ExecutePriceChangeRuleCollectionRequest) error {
	logEntry.WithField("req", req).Debug("about to execute price change Rule Collection")

	uuid, err := domain.ParseUUID(req.Id)
	if err != nil {
		logEntry.WithError(err).Warning("unable to convert to UUID")
		return usecase.ErrDomainInternal
	}

	// Load price change rule collections, Test that it (only) is one
	pricechangerulecollection, err := f.store.PriceChangeRuleCollection().FindByIDs(uuid)
	if err != nil || len(pricechangerulecollection) == 0 {
		logEntry.WithError(err).Warning("unable to retrieve database information")
		return usecase.ErrDatabaseNotFound
	}
	if len(pricechangerulecollection) != 1 {
		logEntry.WithError(err).Warning("Unexpected amount of Price Change Rule Collections with the same ID")
		return usecase.ErrDatabaseInternal
	}

	// Load price change rules
	// TODO: Order by priority in price change rule collection
	priceChangeRules, err := f.store.PriceChangeRule().Load(9999999, 0)
	if err != nil {
		return usecase.ErrDatabaseInternal
	}

	for _, pcr := range priceChangeRules {
		// Load a slice of the affected contracts
		// TODO: Iterate in 100er batches

		filter := map[string]interface{}{
			"productName": pcr.ValidForProductName,
			"commodity":   pcr.ValidForCommodity,
			// TODO: "inArea":
			"status": "active", // TODO: Make dynamicially adjustable
			// TODO: Add time related exclusion criteria
		}

		// Load contractInformations
		contractInformations, err := f.store.ContractInformation().Find(filter, 999999, 0)
		if err != nil {
			logEntry.WithError(err).Warning("unable to retrieve database information")
			return usecase.ErrDatabaseInternal
		}
		/*
			// Now apply the priceChangeRule to each contractInformation
			for _, ci := range contractInformations {
				// First check if a priceChangeOrder is available for the contractInformation and the priceChangeRuleCollection
				priceChangeOrderFilter := map[string]interface{}{
					"priceChangeRuleCollectionId": pcr.ValidForProductName,
					"contractInformationId":       pcr.ValidForCommodity,
				}
				priceChangeOrders, err := f.store.PriceChangeOrder().Find(priceChangeOrderFilter, 999999, 0)
				// Check that it only is one
				// If yes: adjust the existing priceChangeOrder
				// If no: create a new priceChangeOrder

				// Check for Exclusion reasons
				if input.ExcludeSOS == true {
					// TODO: Parse Dates
					if c.StartDate < input.ExcludeSOSFrom {
						c.NewPriceInclude = false
						continue
					}
				}
				// Check for Contract Lifetime
				if input.ExcludeContractduration == true {
					contractStartDateDate, err := time.Parse("2006-01-02", c.StartDate)
					if err != nil {
						fmt.Println("Fehler bei der Datumsberechnung: ", err)
					}
					today := time.Now().Local().AddDate(0, 2, 0)
					contractDays := int(today.Sub(contractStartDateDate).Hours() / 24)
					if input.ExcludeContractdurationDays > contractDays {
						c.NewPriceInclude = false
						continue
					}
				}
				// Add termination date comparison
				// TODO
				// Add Product Change date comparison
				// TODO
				switch input.Typeofchange {
				case "price":
					switch input.Change {
					case "take":
						c.NewPriceInclude = true
						c.NewPriceBase = c.BaseNewPriceProposed
						c.NewPriceKwh = c.KwhNewPriceProposed
					case "set":
						c.NewPriceInclude = true
						c.NewPriceBase = input.Changebase
						c.NewPriceKwh = input.Changekwh
					case "add":
						c.NewPriceInclude = true
						c.NewPriceBase = c.NewPriceBase + input.Changebase
						c.NewPriceKwh = c.NewPriceKwh + input.Changekwh
					case "exclude":
						c.NewPriceInclude = false
					default:
						app.serverErrorResponse(w, r, nil)
						return
					}

					// Check for the Limits
					if input.LimitToCatalogueprice == true {
						if c.CatalogKwhPrice < c.NewPriceKwh {
							c.NewPriceKwh = c.CatalogKwhPrice
						}
						if c.CatalogBasePrice < c.NewPriceBase {
							c.NewPriceBase = c.CatalogBasePrice
						}
					}
					if input.LimitToMax == true {
						if c.NewPriceKwh > input.LimitToMaxKwh {
							c.NewPriceKwh = input.LimitToMaxKwh
						}
						if c.NewPriceBase > input.LimitToMaxBase {
							c.NewPriceBase = input.LimitToMaxBase
						}
					}
					if input.LimitToMin == true {
						if c.NewPriceKwh < input.LimitToMinKwh {
							c.NewPriceKwh = input.LimitToMinKwh
						}
						if c.NewPriceBase < input.LimitToMinBase {
							c.NewPriceBase = input.LimitToMinBase
						}
					}
					if input.LimitToFactor == true {
						// TODO: Check
						if c.NewPriceKwh > (c.CurrentKwhPriceNet * float64(input.LimitToFactorMax/100)) {
							c.NewPriceKwh = c.CurrentKwhPriceNet * float64(input.LimitToFactorMax/100)
						}
						if c.NewPriceBase > (c.CurrentBasePriceNet * float64(input.LimitToFactorMax/100)) {
							c.NewPriceBase = c.CurrentBasePriceNet * float64(input.LimitToFactorMax/100)
						}
						if c.NewPriceKwh < (c.CurrentKwhPriceNet * float64(input.LimitToFactorMin/100)) {
							c.NewPriceKwh = c.CurrentKwhPriceNet * float64(input.LimitToFactorMin/100)
						}
						if c.NewPriceBase < (c.CurrentBasePriceNet * float64(input.LimitToFactorMin/100)) {
							c.NewPriceBase = c.CurrentBasePriceNet * float64(input.LimitToFactorMin/100)
						}
					}
				case "date":
					switch input.Change {
					case "take":
						c.NewPriceStartdate = time.Now().Local().AddDate(0, 2, 0).Format("2006-01-02")
					case "set":
						c.NewPriceStartdate = input.Changedate
					default:
						app.serverErrorResponse(w, r, nil)
						return
					}
				case "communication":
					switch input.Channel {
					case "postmail":
						c.CommunicationChannel = "postmail"
						c.CalculateCommunicationDates(input.Allatonce, input.AllatonceDate, input.Beforechange, input.BeforechangeDays)
					case "email":
						c.CommunicationChannel = "email"
						c.CalculateCommunicationDates(input.Allatonce, input.AllatonceDate, input.Beforechange, input.BeforechangeDays)
					case "both":
						c.CommunicationChannel = "both"
						c.CalculateCommunicationDates(input.Allatonce, input.AllatonceDate, input.Beforechange, input.BeforechangeDays)
					case "none":
						c.CommunicationChannel = ""
						c.CommunicationDate1 = ""
						c.CommunicationDate2 = ""
					default:
						app.serverErrorResponse(w, r, nil)
						return
					}
				default:
					app.serverErrorResponse(w, r, nil)
					return
				}

			}

			// Iterate through the contractInformations
		*/
		fmt.Println("HERE COMES THE CODE.........", contractInformations)
		if err != nil {
			logEntry.WithError(err).Warning("unable to create new price change execution")
			return usecase.ErrDomainInternal
		}
	}

	return nil
}
