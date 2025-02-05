select 
	i.id,
	i."name",
	id.id,
	id.cpu_model,
	id.ram,
	id.year,
	id.screen_size,
	id.capacity,
	id.color,
	id.price 
from pari.items i
left join pari.item_details id on id.item_id = i.id
order by id.created_at desc
limit $1
offset $2;