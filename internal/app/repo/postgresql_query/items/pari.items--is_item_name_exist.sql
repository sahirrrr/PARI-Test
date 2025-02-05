SELECT EXISTS (
    SELECT 1 
    FROM pari.items i
    WHERE i.name = $1
);
