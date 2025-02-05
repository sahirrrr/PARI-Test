update
    pari.item_details
set
    cpu_model = COALESCE($2, cpu_model),
    ram = COALESCE($3, ram),
    year = COALESCE($4, year),
    screen_size = COALESCE($5, screen_size),
    capacity = COALESCE($6, capacity),
    color = COALESCE($7, color),
    price = COALESCE($8, price),
    modified_by = 'SYSTEM',
    modified_at = NOW()
where
    id = $1;