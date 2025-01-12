# Implemented Cases:

## seperate entity for coupon - bxgy

- bxgy coupon is not tighty coupled with items.
- Implemented seperate table for bxgy item which is related to coupon entity (1-1).
- bxgy can be added/updated seperately and can be used further with other discount stregey like. discount on gy items.
- Added transaction (atomacity) for deletion of existing coupon & bxgy_items if exist to maintain the consistancy.

# UnImplemented Cases:

- Min cart limit can ba added
- Max discout amount can be there
- Coupon can have a type like fix, percantage

# Limitations:

# Assumptions:
