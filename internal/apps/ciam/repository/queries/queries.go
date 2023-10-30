package queries

const (
	CreateStore = `INSERT INTO store (
                   id,
                   name,
                   created_at
                   ) 
				    VALUES (
				            :id, 
				            :name, 
				            NOW()
				            ) 
				            RETURNING 
				            id, 
				            name
				    `
)
