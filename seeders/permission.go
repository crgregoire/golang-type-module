package seeders

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/types"
)

var permissions = types.Permissions{
	{
		ID:      uuid.FromStringOrNil("002bff9a-c967-4999-a6b8-99528272f7dc"),
		Name:    "accounts_reader",
		Slug:    "/accounts",
		Actions: []byte("[\"GET\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("009bdebe-e172-446b-b67f-ea3068d29948"),
		Name:    "accounts_admin",
		Slug:    "/accounts",
		Actions: []byte("[\"GET\",\"POST\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("082c4851-a6ac-48b7-ae28-304bf79c9466"),
		Name:    "account_admin",
		Slug:    "/accounts/{account_id}",
		Actions: []byte("[\"GET\",\"PUT\",\"DELETE\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("09dc7b91-e6bd-422d-a81b-9bfd1ac4850b"),
		Name:    "account_reader",
		Slug:    "/accounts/{account_id}",
		Actions: []byte("[\"GET\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("0d88e63f-45d7-42bc-8c59-57d8ac2d17d9"),
		Name:    "account_writer",
		Slug:    "/accounts/{account_id}",
		Actions: []byte("[\"GET\",\"PUT\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("118fea00-a4b8-4d8f-965c-7bad0fc521da"),
		Name:    "account_users_reader",
		Slug:    "/accounts/{account_id}/users",
		Actions: []byte("[\"GET\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("1234a7f2-47a7-4c98-9d6a-1977a55557b4"),
		Name:    "account_users_admin",
		Slug:    "/accounts/{account_id}/users",
		Actions: []byte("[\"GET\",\"POST\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("23b091df-2d15-4e08-bcb9-3eb07d3e28cf"),
		Name:    "account_user_reader",
		Slug:    "/accounts/{account_id}/users/{user_id}",
		Actions: []byte("[\"GET\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("2602bb39-a9fb-459f-b536-aea7fc21c256"),
		Name:    "account_user_writer",
		Slug:    "/accounts/{account_id}/users/{user_id}",
		Actions: []byte("[\"GET\",\"PUT\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("29de0391-9cfb-4bc6-9d2a-142c0fad0df9"),
		Name:    "account_user_admin",
		Slug:    "/accounts/{account_id}/users/{user_id}",
		Actions: []byte("[\"GET\",\"PUT\",\"DELETE\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("9d13abac-7be7-4587-8bf8-caa0c303cfdc"),
		Name:    "account_dispenser_admin",
		Slug:    "/account/dispensers",
		Actions: []byte("[\"GET\",\"PUT\",\"DELETE\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("2ad8d005-419d-4769-b44e-7dbcf888946d"),
		Name:    "pods_reader",
		Slug:    "/pods",
		Actions: []byte("[\"GET\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("447013e5-31af-45b3-b095-8a5e5fcd1b75"),
		Name:    "pods_admin",
		Slug:    "/pods",
		Actions: []byte("[\"GET\",\"POST\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("4ad5db83-22ed-4fc0-a2ae-9378ee05e662"),
		Name:    "pod_reader",
		Slug:    "/pods/{pod_id}",
		Actions: []byte("[\"GET\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("4e2dc0b0-ab31-4fa1-b842-e27c50133bd1"),
		Name:    "pod_admin",
		Slug:    "/pods/{pod_id}",
		Actions: []byte("[\"GET\",\"PUT\",\"DELETE\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("633ee4ed-ec4e-4fd2-b7c2-678a5121a5f0"),
		Name:    "dispensers_reader",
		Slug:    "/dispensers",
		Actions: []byte("[\"GET\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("698405e3-533a-48b5-afc4-d731522ac7bd"),
		Name:    "dispensers_admin",
		Slug:    "/dispensers",
		Actions: []byte("[\"GET\",\"POST\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("6d0c5b7d-2a7c-431e-8282-3ec38a3bcfab"),
		Name:    "dispenser_reader",
		Slug:    "/dispensers/{dispenser_id}",
		Actions: []byte("[\"GET\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("6f15c070-9f64-4b6e-9aef-50c40a0362b3"),
		Name:    "dispenser_admin",
		Slug:    "/dispensers/{dispenser_id}",
		Actions: []byte("[\"GET\",\"PUT\",\"DELETE\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("753220ee-cb62-4023-90c6-f5cbc4218787"),
		Name:    "connections_reader",
		Slug:    "/connections",
		Actions: []byte("[\"GET\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("8a4c1950-df48-4db6-a18e-2c0706012d80"),
		Name:    "connections_admin",
		Slug:    "/connections",
		Actions: []byte("[\"GET\",\"POST\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("93325e97-a58b-4c4e-82f9-349e3b59efaa"),
		Name:    "connection_reader",
		Slug:    "/connections/{connection_id}",
		Actions: []byte("[\"GET\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("985110dd-3295-4533-8a46-2407be51ebf8"),
		Name:    "connection_admin",
		Slug:    "/connections/{connection_id}",
		Actions: []byte("[\"GET\",\"PUT\",\"DELETE\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("99a0759e-18e8-4ac9-9fbb-66f505328a50"),
		Name:    "insertions_reader",
		Slug:    "/insertions",
		Actions: []byte("[\"GET\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("9fef5712-ec5a-486a-8e79-6bf296ae3384"),
		Name:    "insertions_admin",
		Slug:    "/insertions",
		Actions: []byte("[\"GET\",\"POST\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("a0d87693-49ba-4851-a979-756007a6b150"),
		Name:    "insertion_reader",
		Slug:    "/insertions/{insertion_id}",
		Actions: []byte("[\"GET\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("a169ed37-ef4a-408a-9a92-3e58a9939d1d"),
		Name:    "insertion_admin",
		Slug:    "/insertions/{insertion_id}",
		Actions: []byte("[\"GET\",\"PUT\",\"DELETE\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("a42d0d04-9492-4548-a590-e089ae9a4ada"),
		Name:    "regimens_reader",
		Slug:    "/regimens",
		Actions: []byte("[\"GET\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("a42d0d04-9492-4548-a590-e709ae9a4ada"),
		Name:    "regimens_admin",
		Slug:    "/regimens",
		Actions: []byte("[\"GET\",\"PUT\",\"POST\",\"DELETE\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("a8625008-f371-4e40-ae22-368b236323da"),
		Name:    "regimen_reader",
		Slug:    "/regimens/{regimen_id}",
		Actions: []byte("[\"GET\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("abd4be0c-fcb3-4e46-aecd-c76478313284"),
		Name:    "regimen_admin",
		Slug:    "/regimens/{regimen_id}",
		Actions: []byte("[\"GET\",\"PUT\",\"DELETE\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("bbdfc683-46b6-4c61-84d0-4cc1173549b9"),
		Name:    "usages_reader",
		Slug:    "/usages",
		Actions: []byte("[\"GET\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("bd57a2f0-d28e-4f19-9e67-1702176c95e1"),
		Name:    "usages_admin",
		Slug:    "/usages",
		Actions: []byte("[\"GET\",\"POST\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("c13f1669-4d14-448e-995e-4d7e4f3fdeb8"),
		Name:    "usage_reader",
		Slug:    "/usages/{usage_id}",
		Actions: []byte("[\"GET\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("c406d4db-684a-4303-b543-c5d408d35f56"),
		Name:    "usage_admin",
		Slug:    "/usages/{usage_id}",
		Actions: []byte("[\"GET\",\"PUT\",\"DELETE\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("c8c1ac99-9d16-4d4a-bcca-6f9b1e125971"),
		Name:    "permissions_reader",
		Slug:    "/permissions",
		Actions: []byte("[\"GET\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("c9fc605c-f8d2-4446-9bf4-c044ed48d049"),
		Name:    "permissions_reader",
		Slug:    "/permissions",
		Actions: []byte("[\"GET\",\"POST\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("dc8a8a66-0c89-4fa5-99ef-eedeb6abc734"),
		Name:    "permission_reader",
		Slug:    "/permissions/{permission_id}",
		Actions: []byte("[\"GET\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("df009798-f660-4399-9d03-46441c5779ae"),
		Name:    "permission_admin",
		Slug:    "/permissions/{permission_id}",
		Actions: []byte("[\"GET\",\"PUT\",\"DELETE\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("dfb399d1-1238-49cc-8f60-880846b9a93b"),
		Name:    "roles_reader",
		Slug:    "/roles",
		Actions: []byte("[\"GET\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("ee754e19-1715-45a5-ae47-8def127a035e"),
		Name:    "roles_admin",
		Slug:    "/roles",
		Actions: []byte("[\"GET\",\"PUT\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("f033f673-beff-4e36-8fcc-5413e6d0829f"),
		Name:    "role_reader",
		Slug:    "/roles/{role_id}",
		Actions: []byte("[\"GET\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("f4e47676-1178-42ed-9927-6177098595c7"),
		Name:    "role_admin",
		Slug:    "/roles/{role_id}",
		Actions: []byte("[\"GET\",\"PUT\",\"DELETE\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("fd98d07e-2140-4f45-bc98-ca4e948992e0"),
		Name:    "role_permission_admin",
		Slug:    "/roles/{role_id}/{permission_id}",
		Actions: []byte("[\"GET\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("6fa0cfce-1a0f-4b7d-805c-00f5f8666d5b"),
		Name:    "lambda_dispenser_inserted",
		Slug:    "/dispenser/inserted",
		Actions: []byte("[\"POST\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("3af3b1c4-3eba-4229-a5bf-265814444f81"),
		Name:    "lambda_dispenser_dispensed",
		Slug:    "/dispenser/dispensed",
		Actions: []byte("[\"POST\"]"),
	},
	{
		ID:      uuid.FromStringOrNil("98c16e60-9b2a-45ff-a740-30243bc95ba6"),
		Name:    "wp_create_user",
		Slug:    "/wp/createUser",
		Actions: []byte("[\"POST\"]"),
	},
}

func seedPermissions(db *gorm.DB) error {
	if !db.HasTable(&types.Permission{}) {
		if err := db.AutoMigrate(&types.Permission{}).Error; err != nil {
			return err
		}
	}
	for _, permission := range permissions {
		if err := permission.Create(db); err != nil {
			return err
		}
	}
	return nil
}
