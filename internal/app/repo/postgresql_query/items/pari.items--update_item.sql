update 
    pari.items
set
    name = $2,
    modified_by = 'SYSTEM',
    modified_at = NOW()
Where
    id = $1;