-- +migrate Down
DROP TABLE customers;
DROP TABLE price_adjustment_projects;
DROP TABLE price_change_rule_collections;
DROP TABLE price_change_rules;
DROP TABLE contract_informations;
DROP TABLE price_change_orders;
DROP TABLE price_change_executions;

-- +migrate Up
CREATE TABLE customers (
    id          UUID PRIMARY KEY,
    first_name  VARCHAR(255) NOT NULL,
    last_name   VARCHAR(255) NOT NULL,
    email       VARCHAR(255) NOT NULL UNIQUE,
    created_at  TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    updated_at  TIMESTAMPTZ DEFAULT NOW() NOT NULL
);


CREATE TABLE price_adjustment_projects (
    id               UUID PRIMARY KEY,
    name             VARCHAR(255) NOT NULL,
    comment          VARCHAR(255),
    creator          VARCHAR(255),
    creationTime     TIMESTAMPTZ,
    confirmer        VARCHAR(255),
    confirmationTime TIMESTAMPTZ,
    updatedTime      TIMESTAMPTZ,
    executor         VARCHAR(255),
    executionTime    TIMESTAMPTZ,
    locked           BOOL
);

CREATE TABLE price_change_rule_collections(
    id UUID PRIMARY KEY,
    price_adjustment_project_id VARCHAR(255) NOT NULL,
    created_at  TIMESTAMPTZ
);

CREATE TABLE price_change_rules
(
    id                         UUID PRIMARY KEY,
    price_change_rule_collection_id  UUID,

    restore_margin_at_signup      BOOL,
    change_base_price_net_to_amount NUMERIC(8, 2),
    change_kwh_price_net_to_amount  NUMERIC(8, 2),
    change_base_price_net_by_amount NUMERIC(8, 2),
    change_kwh_price_net_by_amount  NUMERIC(8, 2),
    change_base_price_net_by_factor NUMERIC(8, 2),
    change_kwh_price_net_by_factor  NUMERIC(8, 2),

    valid_for_product_name       VARCHAR(1024),
    valid_for_commodity          VARCHAR(255),
    exclude_order_date_from       TIMESTAMPTZ,
    exclude_start_date_from       TIMESTAMPTZ,
    exclude_end_date_until        TIMESTAMPTZ,
    exclude_last_price_change_since TIMESTAMPTZ,

    limit_to_catalogue_price_net        BOOL,
    limit_to_upper_base_price_net        NUMERIC(8, 2),
    limit_to_upper_kwh_price_net         NUMERIC(8, 2),
    limit_to_lower_base_price_net        NUMERIC(8, 2),
    limit_to_lower_kwh_price_net         NUMERIC(8, 2),
    limit_to_max_change_base_price_net    NUMERIC(8, 2),
    limit_to_max_change_kwh_price_net     NUMERIC(8, 2),
    limit_to_min_change_base_price_net    NUMERIC(8, 2),
    limit_to_min_change_kwh_price_net     NUMERIC(8, 2),
    order_in_price_change_rule_collection INT,

        price_adjustment_project_id VARCHAR(255),
    created_at  TIMESTAMPTZ

);


CREATE TABLE  contract_informations (
    id                  UUID PRIMARY KEY,
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
    signup_base_margin    NUMERIC(8, 2),
    signup_kwh_margin     NUMERIC(8, 2),
    current_base_costs    NUMERIC(8, 2),
    current_kwh_costs     NUMERIC(8, 2),
    current_base_margin   NUMERIC(8, 2),
    current_kwh_margin    NUMERIC(8, 2),
    current_base_price_net NUMERIC(8, 2),
    current_kwh_price_net  NUMERIC(8, 2),
    annual_consumption   NUMERIC(12, 4)
);

CREATE TABLE price_change_orders (
    id                      UUID PRIMARY KEY,
    created_at               TIMESTAMPTZ,
    price_change_rule_collection_id         UUID,
    contract_information_id    UUID NOT NULL,

    product_serial_number     VARCHAR(255),
    status                  VARCHAR(255),

    price_valid_since         TIMESTAMPTZ,
    current_base_costs        NUMERIC(8, 2),
    current_kwh_costs         NUMERIC(8, 2),
    current_base_margin       NUMERIC(8, 2),
    current_kwh_margin        NUMERIC(8, 2),
    current_base_price_net     NUMERIC(8, 2),
    current_kwh_price_net      NUMERIC(8, 2),
    annual_consumption       NUMERIC(12, 4),

    price_valid_as_of          TIMESTAMPTZ,
    future_base_costs         NUMERIC(8, 2),
    future_kwh_costs          NUMERIC(8, 2),
    future_kwh_margin         NUMERIC(8, 2),
    future_base_margin        NUMERIC(8, 2),
    future_base_price_net      NUMERIC(8, 2),
    future_kwh_price_net       NUMERIC(8, 2),
    agent_hint_flag           BOOL,
    agent_hint_text           VARCHAR(255),
    communication_flag       BOOL,
    communication_time        TIMESTAMPTZ
);

CREATE TABLE price_change_executions (
    id                      UUID PRIMARY KEY,
    created_at              TIMESTAMPTZ,
    product_serial_Number     VARCHAR(255),
    price_change_order  VARCHAR(255),

    status                  VARCHAR(255),
    execution_time           TIMESTAMPTZ,
    pricechanger_response    VARCHAR(1024),

    price_valid_as_of          TIMESTAMPTZ,
    current_base_price_net     NUMERIC(8, 2),
    future_base_price_net      NUMERIC(8, 2),
    current_kwh_price_net      NUMERIC(8, 2),
    future_kwh_price_net       NUMERIC(8, 2),
    agent_hint_flag           BOOL,
    agent_hint_text           VARCHAR(255),
    communication_flag       BOOL,
    communication_time        TIMESTAMPTZ,
    annual_consumption       NUMERIC(12, 4)
);

