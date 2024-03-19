-- Authentication & Authorization Service
CREATE TABLE users
(
    user_id    UUID PRIMARY KEY,
    username   VARCHAR(50) UNIQUE  NOT NULL,
    password   VARCHAR(255)        NOT NULL,
    email      VARCHAR(100) UNIQUE NOT NULL,
    role       VARCHAR(20)         NOT NULL,
    is_deleted BOOLEAN   DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- User Service
CREATE TABLE user_profiles
(
    user_profile_id UUID PRIMARY KEY,
    user_id         UUID UNIQUE NOT NULL,
    full_name       VARCHAR(100),
    address         TEXT,
    city            VARCHAR(50),
    state           VARCHAR(50),
    country         VARCHAR(50),
    postal_code     VARCHAR(20),
    phone           VARCHAR(20),
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);

CREATE TABLE categories
(
    category_id        UUID PRIMARY KEY,
    name               VARCHAR(100) NOT NULL,
);

-- Product Service
CREATE TABLE products
(
    product_id  UUID PRIMARY KEY,
    name        VARCHAR(100)   NOT NULL,
    description TEXT,
    price       NUMERIC(10, 2) NOT NULL,
    category_id UUID           NOT NULL,
    is_deleted BOOLEAN DEFAULT FALSE,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES categories (category_id)
);

-- Order Service
CREATE TABLE orders
(
    order_id     UUID PRIMARY KEY,
    user_id      UUID           NOT NULL,
    total_amount NUMERIC(10, 2) NOT NULL,
    status       VARCHAR(20)    NOT NULL,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);

CREATE TABLE order_items
(
    order_item_id UUID PRIMARY KEY,
    order_id      UUID           NOT NULL,
    product_id    UUID           NOT NULL,
    quantity      INT            NOT NULL,
    price         NUMERIC(10, 2) NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders (order_id),
    FOREIGN KEY (product_id) REFERENCES products (product_id)
);

-- Payment Service
CREATE TABLE payments
(
    payment_id     UUID PRIMARY KEY,
    order_id       UUID           NOT NULL,
    amount         NUMERIC(10, 2) NOT NULL,
    payment_method VARCHAR(50)    NOT NULL,
    transaction_id VARCHAR(255)   NOT NULL,
    status         VARCHAR(20)    NOT NULL,
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders (order_id)
);

-- Inventory Service
CREATE TABLE inventory
(
    product_id     UUID PRIMARY KEY,
    stock_quantity INT NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products (product_id)
);

-- Shipping Service
CREATE TABLE shipments
(
    shipment_id      UUID PRIMARY KEY,
    order_id         UUID           NOT NULL,
    tracking_number  VARCHAR(100)   NOT NULL,
    shipping_carrier VARCHAR(50)    NOT NULL,
    shipping_cost    NUMERIC(10, 2) NOT NULL,
    shipping_address TEXT           NOT NULL,
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders (order_id)
);

-- Reviews & Ratings Service
CREATE TABLE product_reviews
(
    review_id  UUID PRIMARY KEY,
    product_id UUID NOT NULL,
    user_id    UUID NOT NULL,
    rating     INT  NOT NULL,
    review     TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products (product_id),
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);