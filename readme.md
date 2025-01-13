# Implemented Cases:

- bxgy coupon is loosly coupled with items.
- Implemented seperate table for bxgy item which is related to coupon entity (1-1).
- bxgy can be added/updated seperately.
- Added transaction (atomacity) for deletion of existing coupon & bxgy_items if exist to maintain the consistancy.

# UnImplemented Cases:

- Min cart limit can ba added
- Max discout amount can be there for each coupon
- Coupon can have a type like fix, percantage.
- GY items can be used further with other discount stretgy like. discount on gy items.
- Error handling and Logging can be implemented in much better way, which will help for debugging.

# Limitations:

- Validation function only checks for data and bxgy existance.
- Logging could be iproved to provid more information.
- For large numbers of coupons or cart items, performance optimiztion techniques might be necessary.
- Routes are not sceured.

# Assumptions:

- Postgres service is running on default PORT.
- Go 1.23.2 is installed.
- To keep track of coupon, BxGy Items and coupon usage, I used 3 tables. (coupon, bxgy_items, coupon_usage).
- Currently we only have three types of coupon (CART_WISE, PRODUCT_WISE, BXGY).
- The type of coupon, bxgy and cart in request body is always correct.

## Setup

Run the following command to sync all dependencies:

```bash
go mod tidy
```

## API Routes

### 1. Register a Coupon

Registers a new coupon.

- **Endpoint:** `POST /coupons`
- **API URL:** `localhost:8080/coupons`

**Example Request Body:**

```json
{
  "code": "NEWYEAR",
  "discountType": 2,
  "discountValue": 20,
  "startDate": "2024-02-15T00:00:00Z",
  "endDate": "2025-03-15T00:00:00Z",
  "details": { "disc": "use bxgy_items table for more details" },
  "bxItemList": ["SKU3", "SKU4"], // Optional
  "gyItemList": ["SKU1", "SKU2", "SKU3"] // Optional
}
```

### 2. Get all Coupons

To fetch all the existing coupons.

- **Endpoint:** `GET /coupons`
- **API URL:** `localhost:8080/coupons`

### 3. Get a Coupon By its ID

To fetch a coupon by coupon id.

- **Endpoint:** `GET /coupons/:id`
- **API URL:** `localhost:8080/coupons/:id`

### 4. Update a Coupon By its ID

To update a coupon by coupon id.

- **Endpoint:** `PUT /coupons/:id`
- **API URL:** `localhost:8080/coupons/:id`

### 5. Delete a Coupon By its ID

To delete a coupon by coupon id.

- **Endpoint:** `PUT /coupons/:id`
- **API URL:** `localhost:8080/coupons/:id`

### 6. Get List of Coupons Applicable to Current Cart Items

To fetch all coupons that can be applied on current cart items.

- **Endpoint:** `POST /applicable-coupons`
- **API URL:** `localhost:8080/applicable-coupons`

**Example Request Body:**

```json
{
  "cart": {
    "items": [
      { "sku": "SKU1", "quantity": 1, "price": 99.0 },
      { "sku": "SKU2", "quantity": 2, "price": 199.0 },
      { "sku": "SKU3", "quantity": 2, "price": 299.0 }
    ],
    "cartTotal": 595.0
  }
}
```

### 7. Apply a Coupon to Current Cart Items and Get Discounted Values.

To get the updated cart total and dicounted values according to coupon applied.

- **Endpoint:** `POST /apply-coupon/:id`
- **API URL:** `localhost:8080/apply-coupon/:id`

**Example Request Body:**

```json
{
  "cart": {
    "items": [
      { "sku": "SKU1", "quantity": 1, "price": 99.0 },
      { "sku": "SKU2", "quantity": 2, "price": 199.0 },
      { "sku": "SKU3", "quantity": 2, "price": 299.0 }
    ],
    "cartTotal": 595.0,
    "appliedCoupons": [2]
  }
}
```
