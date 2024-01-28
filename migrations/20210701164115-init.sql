
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


CREATE TABLE price_adjustment_projects (
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
    locked           BOOL,
    PRIMARY KEY (id)
);

CREATE TABLE price_change_rule_collections(
    id UUID NOT NULL,
    priceAdjustmentProject UUID NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE price_change_rules
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
    orderInPriceChangeRuleCollection INT,
    PRIMARY KEY (id)
);


CREATE TABLE  contract_informations (
    id                  UUID NOT NULL,
    snapshot_time        TIMESTAMPTZ,

    mba                 VARCHAR(255),
    product_serial_number VARCHAR(255),
    product_name         VARCHAR(255),
    in_area              BOOL,
    commodity           VARCHAR(255),

    order_date           TIMESTAMPTZ,
    start_date           TIMESTAMPTZ,
    end_date             TIMESTAMPTZ,
    status              VARCHAR(255),
    price_guarantee_until TIMESTAMPTZ,
    price_change_planned  BOOL,

    price_valid_since     TIMESTAMPTZ,
    current_base_costs    NUMERIC(6, 2),
    current_kwh_costs     NUMERIC(6, 2),
    current_base_margin   NUMERIC(6, 2),
    current_kwh_margin    NUMERIC(6, 2),
    current_base_price_net NUMERIC(6, 2),
    current_kwh_price_net  NUMERIC(6, 2),
    annual_consumption   NUMERIC(12, 4),
    PRIMARY KEY (id)
);

CREATE TABLE price_change_orders (
    id                      UUID,
    created_at               TIMESTAMPTZ,
    price_change_rule         VARCHAR(255),

    product_serial_number     VARCHAR(255),
    status                  VARCHAR(255),

    price_valid_since         TIMESTAMPTZ,
    current_base_costs        NUMERIC(6, 2),
    current_kwh_costs         NUMERIC(6, 2),
    current_base_margin       NUMERIC(6, 2),
    current_kwh_margin        NUMERIC(6, 2),
    current_base_price_net     NUMERIC(6, 2),
    current_kwh_price_net      NUMERIC(6, 2),
    annual_consumption       NUMERIC(12, 4),

    price_valid_as_of          TIMESTAMPTZ,
    future_base_costs         NUMERIC(6, 2),
    future_kwh_costs          NUMERIC(6, 2),
    future_kwh_margin         NUMERIC(6, 2),
    future_base_margin        NUMERIC(6, 2),
    future_base_price_net      NUMERIC(6, 2),
    future_kwh_price_net       NUMERIC(6, 2),
    agent_hint_flag           BOOL,
    agent_hint_text           VARCHAR(255),
    communication_flag       BOOL,
    communiction_time        TIMESTAMPTZ,
    PRIMARY KEY (id)
);

CREATE TABLE price_change_executions (
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
    annualConsumption       NUMERIC(12, 4),
    PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE customers;
DROP TABLE price_adjustment_projects;
DROP TABLE price_change_rule_collections;
DROP TABLE price_change_rules;
DROP TABLE contract_informations;
DROP TABLE price_change_orders;
DROP TABLE price_change_executions;