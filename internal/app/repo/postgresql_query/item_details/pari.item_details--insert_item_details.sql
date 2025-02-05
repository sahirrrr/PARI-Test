insert into pari.item_details(item_id, cpu_model, ram, year, screen_size, capacity, color, price)
values($1, $2, $3, $4, $5, $6, $7, $8)
returning id;