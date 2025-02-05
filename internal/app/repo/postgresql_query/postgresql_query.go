package postgresql_query

import _ "embed"

var (
	// -------------------------
	// TABLE ITEMS
	// -------------------------

	//go:embed items/pari.items--get_list_items.sql
	GetListItems string

	//go:embed items/pari.items--get_item_by_id.sql
	GetItemByID string

	//go:embed items/pari.items--is_item_name_exist.sql
	IsItemNameExist string

	//go:embed items/pari.items--insert_item.sql
	InsertItem string

	//go:embed items/pari.items--delete_item.sql
	DeleteItem string

	//go:embed items/pari.items--update_item.sql
	UpdateItem string

	// -------------------------
	// TABLE ITEM_DETAILS
	// -------------------------

	//go:embed item_details/pari.item_details--insert_item_details.sql
	InsertItemDetails string

	//go:embed item_details/pari.item_details--update_item_details.sql
	UpdateItemDetails string
)
