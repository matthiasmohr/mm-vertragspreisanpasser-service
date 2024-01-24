
-- +migrate Up
CREATE TABLE customers (
    id          UUID NOT NULL,
    first_name  VARCHAR(255) NOT NULL,
    last_name   VARCHAR(255) NOT NULL,
    email       VARCHAR(255) NOT NULL UNIQUE,
    created_at  TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    updated_at  TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    PRIMARY KEY (id)
);


CREATE TABLE priceAdjustmentProject (
    id               UUID         NOT NULL,
    name             VARCHAR(255) NOT NULL,
    comment          VARCHAR(255),
    creator          VARCHAR(255),
    creationTime     TIMESTAMPTZ,
    confirmer        VARCHAR(255),
    confirmationTime TIMESTAMPTZ,
    updatedTime      TIMESTAMPTZ,
    executor         VARCHAR(255),
    executionTime    TIMESTAMPTZ,
    PRIMARY KEY (id)
);

CREATE TABLE priceChangeRuleCollection(
    id UUID NOT NULL,
    priceAdjustmentProject UUID NOT NULL
);

CREATE TABLE priceChangeRule
(
    id                         UUID NOT NULL,
    priceChangeRuleCollection  UUID,

    restoreMarginAtSignup      BOOL,
    changeBasePriceNetToAmount NUMERIC(6, 2),
    changeKwhPriceNetToAmount  NUMERIC(6, 2),
    changeBasePriceNetByAmount NUMERIC(6, 2),
    changeKwhPriceNetByAmount  NUMERIC(6, 2),
    changeBasePriceNetByFactor NUMERIC(6, 2),
    changeKwhPriceNetByFactor  NUMERIC(6, 2),

    validForProductNames       VARCHAR(1024),
    validForCommodity          VARCHAR(255),
    excludeOrderDateFrom       TIMESTAMPTZ,
    excludeStartDateFrom       TIMESTAMPTZ,
    excludeEndDateUntil        TIMESTAMPTZ,
    excludeLastPriceChangeSince TIMESTAMPTZ,

    limitToCataloguePriceNet        BOOL,
    limitToUpperBasePriceNet        NUMERIC(6, 2),
    limitToUpperKwhPriceNet         NUMERIC(4, 2),
    limitToLowerBasePriceNet        NUMERIC(6, 2),
    limitToLowerKwhPriceNet         NUMERIC(6, 2),
    limitToMaxChangeBasePriceNet    NUMERIC(6, 2),
    limitToMaxChangeKwhPriceNet     NUMERIC(6, 2),
    limitToMinChangeBasePriceNet    NUMERIC(6, 2),
    limitToMinChangeKwhPriceNet     NUMERIC(6, 2),
    orderInPriceChangeRuleCollection INT
);


CREATE TABLE  contractInformation (
    id                  UUID NOT NULL,
    snapshotTime        TIMESTAMPTZ,

    mba                 VARCHAR(255),
    productSerialNumber VARCHAR(255),
    productName         VARCHAR(255),
    inArea              BOOL,
    commodity           VARCHAR(255),

    orderDate           TIMESTAMPTZ,
    startDate           TIMESTAMPTZ,
    endDate             TIMESTAMPTZ,
    status              VARCHAR(255),
    priceGuaranteeUntil TIMESTAMPTZ,
    priceChangePlanned  BOOL,

    priceValidSince     TIMESTAMPTZ,
    currentBaseCosts    NUMERIC(6, 2),
    currentKwhCosts     NUMERIC(6, 2),
    currentBaseMargin   NUMERIC(6, 2),
    currentKwhMargin    NUMERIC(6, 2),
    currentBasePriceNet NUMERIC(6, 2),
    currentKwhPriceNet  NUMERIC(6, 2),
    annualConsumption   NUMERIC(12, 4)
);

CREATE TABLE priceChangeOrder (
    id                      UUID,
    priceChangeRule         UUID NOT NULL,
    contractInformation     UUID NOT NULL,

    status                  VARCHAR(255),

    priceValidSince         TIMESTAMPTZ,
    currentBaseCosts        NUMERIC(6, 2),
    currentKwhCosts         NUMERIC(6, 2),
    currentBaseMargin       NUMERIC(6, 2),
    currentKwhMargin        NUMERIC(6, 2),
    currentBasePriceNet     NUMERIC(6, 2),
    currentKwhPriceNet      NUMERIC(6, 2),
    annualConsumption       NUMERIC(12, 4),

    priceValidAsOf          TIMESTAMPTZ,
    futureBaseCosts         NUMERIC(6, 2),
    futureKwhCosts          NUMERIC(6, 2),
    futureKwhMargin         NUMERIC(6, 2),
    futureBaseMargin        NUMERIC(6, 2),
    futureBasePriceNet      NUMERIC(6, 2),
    futureKwhPriceNet       NUMERIC(6, 2),
    agentHintFlag           BOOL,
    agentHintText           VARCHAR(255),
    communicationFlag       BOOL,
    communictionTime        TIMESTAMPTZ
);

CREATE TABLE priceChangeExecution (
    id                      UUID,
    productSerialNumber     VARCHAR(255),
    preisanpassungsprojekt  UUID NOT NULL,

    status                  VARCHAR(255),
    executionTime           TIMESTAMPTZ,
    pricechangerResponse    VARCHAR(1024),

    priceValidAsOf          TIMESTAMPTZ,
    currentBasePriceNet     NUMERIC(6, 2),
    futureBasePriceNet      NUMERIC(6, 2),
    currentKwhPriceNet      NUMERIC(6, 2),
    futureKwhPriceNet       NUMERIC(6, 2),
    agentHintFlag           BOOL,
    agentHintText           VARCHAR(255),
    communicationFlag       BOOL,
    communictionTime        TIMESTAMPTZ,
    annualConsumption       NUMERIC(12, 4)
);

-- +migrate Down
DROP TABLE customers;
DROP TABLE priceAdjustmentProject;
DROP TABLE priceChangeRuleCollection;
DROP TABLE priceChangeRule;
DROP TABLE contractInformation;
DROP TABLE priceChangeOrder;
DROP TABLE priceChangeExecution;