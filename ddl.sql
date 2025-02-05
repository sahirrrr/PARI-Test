CREATE TABLE pari.items (
    id SERIAL PRIMARY KEY,           
    name VARCHAR(255) NOT NULL,
    created_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM',      
    created_at TIMESTAMP NOT NULL DEFAULT NOW(), 
    modified_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM', 
    modified_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_by VARCHAR(255) NULL,               
    deleted_at TIMESTAMP NULL                                        
);
CREATE UNIQUE INDEX idx_items_name ON pari.items USING BTREE (name);

CREATE TABLE pari.item_details (
    id SERIAL PRIMARY KEY,           
   	item_id INT NOT NULL REFERENCES pari.items(id) ON DELETE CASCADE,
    color VARCHAR(255) NULL,
    capacity INT NULL,
    screen_size DECIMAL(5,2) NULL,
    price numeric(50, 3) NOT NULL DEFAULT 0,
    year SMALLINT NULL,
    cpu_model VARCHAR(255) NULL,
    ram SMALLINT NULL,
    created_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM',      
    created_at TIMESTAMP NOT NULL DEFAULT NOW(), 
    modified_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM', 
    modified_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_by VARCHAR(255) NULL,               
    deleted_at TIMESTAMP NULL                                        
);