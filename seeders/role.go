package seeders

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/types"
)

var roles = types.Roles{
	{
		ID:   uuid.FromStringOrNil("26841819-b845-4ecb-aa71-ebbbbfc60a6e"),
		Name: "user",
		Meta: nil,
	},
	{
		ID:   uuid.FromStringOrNil("0109f277-66cd-46f0-9107-7aa4ecb2201d"),
		Name: "owner",
		Meta: nil,
	},
	{
		ID:   uuid.FromStringOrNil("6146f7d1-7f7c-4519-a243-dfe71ae2284d"),
		Name: "affiliate",
		Meta: nil,
	},
	{
		ID:   uuid.FromStringOrNil("6d814967-ebdb-4ba7-b8e8-2a94ee5fafeb"),
		Name: "partner",
		Meta: nil,
	},
	{
		ID:   uuid.FromStringOrNil("9651b56b-9781-42a6-99b5-4db352474f2a"),
		Name: "marketing",
		Meta: nil,
	},
	{
		ID:   uuid.FromStringOrNil("ccb900be-f5a8-4a98-a493-00999eadcc42"),
		Name: "customer_service",
		Meta: nil,
	},
	{
		ID:   uuid.FromStringOrNil("d9be45aa-0d38-412e-8eee-cc0fe9b38239"),
		Name: "developer",
		Meta: nil,
	},
	{
		ID:   uuid.FromStringOrNil("f4e17fdb-2728-4b61-af7f-ba220e89fa86"),
		Name: "master",
		Meta: nil,
	},
	{
		ID:   uuid.FromStringOrNil("b22510b1-b501-4c68-802a-e0ebce2b8307"),
		Name: "internal_application",
		Meta: nil,
	},
}

var rolePermissions = map[uuid.UUID][]types.Permission{
	uuid.FromStringOrNil("26841819-b845-4ecb-aa71-ebbbbfc60a6e"): {
		{
			ID: uuid.FromStringOrNil("002bff9a-c967-4999-a6b8-99528272f7dc"),
		},
		{
			ID: uuid.FromStringOrNil("09dc7b91-e6bd-422d-a81b-9bfd1ac4850b"),
		},
		{
			ID: uuid.FromStringOrNil("118fea00-a4b8-4d8f-965c-7bad0fc521da"),
		},
		{
			ID: uuid.FromStringOrNil("23b091df-2d15-4e08-bcb9-3eb07d3e28cf"),
		},
		{
			ID: uuid.FromStringOrNil("2602bb39-a9fb-459f-b536-aea7fc21c256"),
		},
		{
			ID: uuid.FromStringOrNil("2ad8d005-419d-4769-b44e-7dbcf888946d"),
		},
		{
			ID: uuid.FromStringOrNil("4ad5db83-22ed-4fc0-a2ae-9378ee05e662"),
		},
		{
			ID: uuid.FromStringOrNil("633ee4ed-ec4e-4fd2-b7c2-678a5121a5f0"),
		},
		{
			ID: uuid.FromStringOrNil("6d0c5b7d-2a7c-431e-8282-3ec38a3bcfab"),
		},
		{
			ID: uuid.FromStringOrNil("753220ee-cb62-4023-90c6-f5cbc4218787"),
		},
		{
			ID: uuid.FromStringOrNil("93325e97-a58b-4c4e-82f9-349e3b59efaa"),
		},
		{
			ID: uuid.FromStringOrNil("99a0759e-18e8-4ac9-9fbb-66f505328a50"),
		},
		{
			ID: uuid.FromStringOrNil("a0d87693-49ba-4851-a979-756007a6b150"),
		},
		{
			ID: uuid.FromStringOrNil("a42d0d04-9492-4548-a590-e089ae9a4ada"),
		},
		{
			ID: uuid.FromStringOrNil("a8625008-f371-4e40-ae22-368b236323da"),
		},
		{
			ID: uuid.FromStringOrNil("bbdfc683-46b6-4c61-84d0-4cc1173549b9"),
		},
	},
	uuid.FromStringOrNil("0109f277-66cd-46f0-9107-7aa4ecb2201d"): {
		{
			ID: uuid.FromStringOrNil("9d13abac-7be7-4587-8bf8-caa0c303cfdc"),
		},
		{
			ID: uuid.FromStringOrNil("009bdebe-e172-446b-b67f-ea3068d29948"),
		},
		{
			ID: uuid.FromStringOrNil("082c4851-a6ac-48b7-ae28-304bf79c9466"),
		},
		{
			ID: uuid.FromStringOrNil("0d88e63f-45d7-42bc-8c59-57d8ac2d17d9"),
		},
		{
			ID: uuid.FromStringOrNil("1234a7f2-47a7-4c98-9d6a-1977a55557b4"),
		},
		{
			ID: uuid.FromStringOrNil("447013e5-31af-45b3-b095-8a5e5fcd1b75"),
		},
		{
			ID: uuid.FromStringOrNil("4e2dc0b0-ab31-4fa1-b842-e27c50133bd1"),
		},
		{
			ID: uuid.FromStringOrNil("698405e3-533a-48b5-afc4-d731522ac7bd"),
		},
		{
			ID: uuid.FromStringOrNil("6f15c070-9f64-4b6e-9aef-50c40a0362b3"),
		},
		{
			ID: uuid.FromStringOrNil("8a4c1950-df48-4db6-a18e-2c0706012d80"),
		},
		{
			ID: uuid.FromStringOrNil("985110dd-3295-4533-8a46-2407be51ebf8"),
		},
		{
			ID: uuid.FromStringOrNil("99a0759e-18e8-4ac9-9fbb-66f505328a50"),
		},
		{
			ID: uuid.FromStringOrNil("a169ed37-ef4a-408a-9a92-3e58a9939d1d"),
		},
		{
			ID: uuid.FromStringOrNil("a42d0d04-9492-4548-a590-e709ae9a4ada"),
		},
		{
			ID: uuid.FromStringOrNil("abd4be0c-fcb3-4e46-aecd-c76478313284"),
		},
		{
			ID: uuid.FromStringOrNil("c406d4db-684a-4303-b543-c5d408d35f56"),
		},
	},
	uuid.FromStringOrNil("6146f7d1-7f7c-4519-a243-dfe71ae2284d"): {
		{
			ID: uuid.FromStringOrNil("002bff9a-c967-4999-a6b8-99528272f7dc"),
		},
		{
			ID: uuid.FromStringOrNil("09dc7b91-e6bd-422d-a81b-9bfd1ac4850b"),
		},
		{
			ID: uuid.FromStringOrNil("118fea00-a4b8-4d8f-965c-7bad0fc521da"),
		},
		{
			ID: uuid.FromStringOrNil("23b091df-2d15-4e08-bcb9-3eb07d3e28cf"),
		},
		{
			ID: uuid.FromStringOrNil("2602bb39-a9fb-459f-b536-aea7fc21c256"),
		},
		{
			ID: uuid.FromStringOrNil("2ad8d005-419d-4769-b44e-7dbcf888946d"),
		},
		{
			ID: uuid.FromStringOrNil("4ad5db83-22ed-4fc0-a2ae-9378ee05e662"),
		},
		{
			ID: uuid.FromStringOrNil("633ee4ed-ec4e-4fd2-b7c2-678a5121a5f0"),
		},
		{
			ID: uuid.FromStringOrNil("6d0c5b7d-2a7c-431e-8282-3ec38a3bcfab"),
		},
		{
			ID: uuid.FromStringOrNil("753220ee-cb62-4023-90c6-f5cbc4218787"),
		},
		{
			ID: uuid.FromStringOrNil("93325e97-a58b-4c4e-82f9-349e3b59efaa"),
		},
		{
			ID: uuid.FromStringOrNil("99a0759e-18e8-4ac9-9fbb-66f505328a50"),
		},
		{
			ID: uuid.FromStringOrNil("a0d87693-49ba-4851-a979-756007a6b150"),
		},
		{
			ID: uuid.FromStringOrNil("a42d0d04-9492-4548-a590-e089ae9a4ada"),
		},
		{
			ID: uuid.FromStringOrNil("a8625008-f371-4e40-ae22-368b236323da"),
		},
		{
			ID: uuid.FromStringOrNil("bbdfc683-46b6-4c61-84d0-4cc1173549b9"),
		},
		{
			ID: uuid.FromStringOrNil("c13f1669-4d14-448e-995e-4d7e4f3fdeb8"),
		},
	},
	uuid.FromStringOrNil("6d814967-ebdb-4ba7-b8e8-2a94ee5fafeb"): {
		{
			ID: uuid.FromStringOrNil("002bff9a-c967-4999-a6b8-99528272f7dc"),
		},
		{
			ID: uuid.FromStringOrNil("09dc7b91-e6bd-422d-a81b-9bfd1ac4850b"),
		},
		{
			ID: uuid.FromStringOrNil("118fea00-a4b8-4d8f-965c-7bad0fc521da"),
		},
		{
			ID: uuid.FromStringOrNil("23b091df-2d15-4e08-bcb9-3eb07d3e28cf"),
		},
		{
			ID: uuid.FromStringOrNil("2602bb39-a9fb-459f-b536-aea7fc21c256"),
		},
		{
			ID: uuid.FromStringOrNil("2ad8d005-419d-4769-b44e-7dbcf888946d"),
		},
		{
			ID: uuid.FromStringOrNil("4ad5db83-22ed-4fc0-a2ae-9378ee05e662"),
		},
		{
			ID: uuid.FromStringOrNil("633ee4ed-ec4e-4fd2-b7c2-678a5121a5f0"),
		},
		{
			ID: uuid.FromStringOrNil("6d0c5b7d-2a7c-431e-8282-3ec38a3bcfab"),
		},
		{
			ID: uuid.FromStringOrNil("753220ee-cb62-4023-90c6-f5cbc4218787"),
		},
		{
			ID: uuid.FromStringOrNil("93325e97-a58b-4c4e-82f9-349e3b59efaa"),
		},
		{
			ID: uuid.FromStringOrNil("99a0759e-18e8-4ac9-9fbb-66f505328a50"),
		},
		{
			ID: uuid.FromStringOrNil("a0d87693-49ba-4851-a979-756007a6b150"),
		},
		{
			ID: uuid.FromStringOrNil("a42d0d04-9492-4548-a590-e089ae9a4ada"),
		},
		{
			ID: uuid.FromStringOrNil("a8625008-f371-4e40-ae22-368b236323da"),
		},
		{
			ID: uuid.FromStringOrNil("bbdfc683-46b6-4c61-84d0-4cc1173549b9"),
		},
		{
			ID: uuid.FromStringOrNil("c13f1669-4d14-448e-995e-4d7e4f3fdeb8"),
		},
	},
	uuid.FromStringOrNil("9651b56b-9781-42a6-99b5-4db352474f2a"): {
		{
			ID: uuid.FromStringOrNil("002bff9a-c967-4999-a6b8-99528272f7dc"),
		},
		{
			ID: uuid.FromStringOrNil("09dc7b91-e6bd-422d-a81b-9bfd1ac4850b"),
		},
		{
			ID: uuid.FromStringOrNil("118fea00-a4b8-4d8f-965c-7bad0fc521da"),
		},
		{
			ID: uuid.FromStringOrNil("23b091df-2d15-4e08-bcb9-3eb07d3e28cf"),
		},
		{
			ID: uuid.FromStringOrNil("2602bb39-a9fb-459f-b536-aea7fc21c256"),
		},
		{
			ID: uuid.FromStringOrNil("2ad8d005-419d-4769-b44e-7dbcf888946d"),
		},
		{
			ID: uuid.FromStringOrNil("4ad5db83-22ed-4fc0-a2ae-9378ee05e662"),
		},
		{
			ID: uuid.FromStringOrNil("633ee4ed-ec4e-4fd2-b7c2-678a5121a5f0"),
		},
		{
			ID: uuid.FromStringOrNil("6d0c5b7d-2a7c-431e-8282-3ec38a3bcfab"),
		},
		{
			ID: uuid.FromStringOrNil("753220ee-cb62-4023-90c6-f5cbc4218787"),
		},
		{
			ID: uuid.FromStringOrNil("93325e97-a58b-4c4e-82f9-349e3b59efaa"),
		},
		{
			ID: uuid.FromStringOrNil("99a0759e-18e8-4ac9-9fbb-66f505328a50"),
		},
		{
			ID: uuid.FromStringOrNil("a0d87693-49ba-4851-a979-756007a6b150"),
		},
		{
			ID: uuid.FromStringOrNil("a42d0d04-9492-4548-a590-e089ae9a4ada"),
		},
		{
			ID: uuid.FromStringOrNil("a8625008-f371-4e40-ae22-368b236323da"),
		},
		{
			ID: uuid.FromStringOrNil("bbdfc683-46b6-4c61-84d0-4cc1173549b9"),
		},
		{
			ID: uuid.FromStringOrNil("c13f1669-4d14-448e-995e-4d7e4f3fdeb8"),
		},
	},
	uuid.FromStringOrNil("ccb900be-f5a8-4a98-a493-00999eadcc42"): {
		{
			ID: uuid.FromStringOrNil("c9fc605c-f8d2-4446-9bf4-c044ed48d049"),
		},
		{
			ID: uuid.FromStringOrNil("dc8a8a66-0c89-4fa5-99ef-eedeb6abc734"),
		},
		{
			ID: uuid.FromStringOrNil("dfb399d1-1238-49cc-8f60-880846b9a93b"),
		},
		{
			ID: uuid.FromStringOrNil("f033f673-beff-4e36-8fcc-5413e6d0829f"),
		},
		{
			ID: uuid.FromStringOrNil("009bdebe-e172-446b-b67f-ea3068d29948"),
		},
		{
			ID: uuid.FromStringOrNil("082c4851-a6ac-48b7-ae28-304bf79c9466"),
		},
		{
			ID: uuid.FromStringOrNil("0d88e63f-45d7-42bc-8c59-57d8ac2d17d9"),
		},
		{
			ID: uuid.FromStringOrNil("1234a7f2-47a7-4c98-9d6a-1977a55557b4"),
		},
		{
			ID: uuid.FromStringOrNil("447013e5-31af-45b3-b095-8a5e5fcd1b75"),
		},
		{
			ID: uuid.FromStringOrNil("4e2dc0b0-ab31-4fa1-b842-e27c50133bd1"),
		},
		{
			ID: uuid.FromStringOrNil("698405e3-533a-48b5-afc4-d731522ac7bd"),
		},
		{
			ID: uuid.FromStringOrNil("6f15c070-9f64-4b6e-9aef-50c40a0362b3"),
		},
		{
			ID: uuid.FromStringOrNil("8a4c1950-df48-4db6-a18e-2c0706012d80"),
		},
		{
			ID: uuid.FromStringOrNil("985110dd-3295-4533-8a46-2407be51ebf8"),
		},
		{
			ID: uuid.FromStringOrNil("99a0759e-18e8-4ac9-9fbb-66f505328a50"),
		},
		{
			ID: uuid.FromStringOrNil("a169ed37-ef4a-408a-9a92-3e58a9939d1d"),
		},
		{
			ID: uuid.FromStringOrNil("a42d0d04-9492-4548-a590-e709ae9a4ada"),
		},
		{
			ID: uuid.FromStringOrNil("abd4be0c-fcb3-4e46-aecd-c76478313284"),
		},
		{
			ID: uuid.FromStringOrNil("c406d4db-684a-4303-b543-c5d408d35f56"),
		},
	},
	uuid.FromStringOrNil("d9be45aa-0d38-412e-8eee-cc0fe9b38239"): {
		{
			ID: uuid.FromStringOrNil("df009798-f660-4399-9d03-46441c5779ae"),
		},
		{
			ID: uuid.FromStringOrNil("ee754e19-1715-45a5-ae47-8def127a035e"),
		},
		{
			ID: uuid.FromStringOrNil("f4e47676-1178-42ed-9927-6177098595c7"),
		},
		{
			ID: uuid.FromStringOrNil("fd98d07e-2140-4f45-bc98-ca4e948992e0"),
		},
		{
			ID: uuid.FromStringOrNil("009bdebe-e172-446b-b67f-ea3068d29948"),
		},
		{
			ID: uuid.FromStringOrNil("082c4851-a6ac-48b7-ae28-304bf79c9466"),
		},
		{
			ID: uuid.FromStringOrNil("0d88e63f-45d7-42bc-8c59-57d8ac2d17d9"),
		},
		{
			ID: uuid.FromStringOrNil("1234a7f2-47a7-4c98-9d6a-1977a55557b4"),
		},
		{
			ID: uuid.FromStringOrNil("447013e5-31af-45b3-b095-8a5e5fcd1b75"),
		},
		{
			ID: uuid.FromStringOrNil("4e2dc0b0-ab31-4fa1-b842-e27c50133bd1"),
		},
		{
			ID: uuid.FromStringOrNil("698405e3-533a-48b5-afc4-d731522ac7bd"),
		},
		{
			ID: uuid.FromStringOrNil("6f15c070-9f64-4b6e-9aef-50c40a0362b3"),
		},
		{
			ID: uuid.FromStringOrNil("8a4c1950-df48-4db6-a18e-2c0706012d80"),
		},
		{
			ID: uuid.FromStringOrNil("985110dd-3295-4533-8a46-2407be51ebf8"),
		},
		{
			ID: uuid.FromStringOrNil("99a0759e-18e8-4ac9-9fbb-66f505328a50"),
		},
		{
			ID: uuid.FromStringOrNil("a169ed37-ef4a-408a-9a92-3e58a9939d1d"),
		},
		{
			ID: uuid.FromStringOrNil("a42d0d04-9492-4548-a590-e709ae9a4ada"),
		},
		{
			ID: uuid.FromStringOrNil("abd4be0c-fcb3-4e46-aecd-c76478313284"),
		},
		{
			ID: uuid.FromStringOrNil("c406d4db-684a-4303-b543-c5d408d35f56"),
		},
	},
	uuid.FromStringOrNil("f4e17fdb-2728-4b61-af7f-ba220e89fa86"): {
		{
			ID: uuid.FromStringOrNil("9d13abac-7be7-4587-8bf8-caa0c303cfdc"),
		},
		{
			ID: uuid.FromStringOrNil("29de0391-9cfb-4bc6-9d2a-142c0fad0df9"),
		},
		{
			ID: uuid.FromStringOrNil("df009798-f660-4399-9d03-46441c5779ae"),
		},
		{
			ID: uuid.FromStringOrNil("ee754e19-1715-45a5-ae47-8def127a035e"),
		},
		{
			ID: uuid.FromStringOrNil("f4e47676-1178-42ed-9927-6177098595c7"),
		},
		{
			ID: uuid.FromStringOrNil("fd98d07e-2140-4f45-bc98-ca4e948992e0"),
		},
		{
			ID: uuid.FromStringOrNil("009bdebe-e172-446b-b67f-ea3068d29948"),
		},
		{
			ID: uuid.FromStringOrNil("082c4851-a6ac-48b7-ae28-304bf79c9466"),
		},
		{
			ID: uuid.FromStringOrNil("0d88e63f-45d7-42bc-8c59-57d8ac2d17d9"),
		},
		{
			ID: uuid.FromStringOrNil("1234a7f2-47a7-4c98-9d6a-1977a55557b4"),
		},
		{
			ID: uuid.FromStringOrNil("447013e5-31af-45b3-b095-8a5e5fcd1b75"),
		},
		{
			ID: uuid.FromStringOrNil("4e2dc0b0-ab31-4fa1-b842-e27c50133bd1"),
		},
		{
			ID: uuid.FromStringOrNil("698405e3-533a-48b5-afc4-d731522ac7bd"),
		},
		{
			ID: uuid.FromStringOrNil("6f15c070-9f64-4b6e-9aef-50c40a0362b3"),
		},
		{
			ID: uuid.FromStringOrNil("8a4c1950-df48-4db6-a18e-2c0706012d80"),
		},
		{
			ID: uuid.FromStringOrNil("985110dd-3295-4533-8a46-2407be51ebf8"),
		},
		{
			ID: uuid.FromStringOrNil("99a0759e-18e8-4ac9-9fbb-66f505328a50"),
		},
		{
			ID: uuid.FromStringOrNil("a169ed37-ef4a-408a-9a92-3e58a9939d1d"),
		},
		{
			ID: uuid.FromStringOrNil("a42d0d04-9492-4548-a590-e709ae9a4ada"),
		},
		{
			ID: uuid.FromStringOrNil("abd4be0c-fcb3-4e46-aecd-c76478313284"),
		},
		{
			ID: uuid.FromStringOrNil("c406d4db-684a-4303-b543-c5d408d35f56"),
		},
	},
	uuid.FromStringOrNil("b22510b1-b501-4c68-802a-e0ebce2b8307"): {
		{
			ID: uuid.FromStringOrNil("6fa0cfce-1a0f-4b7d-805c-00f5f8666d5b"),
		},
		{
			ID: uuid.FromStringOrNil("3af3b1c4-3eba-4229-a5bf-265814444f81"),
		},
		{
			ID: uuid.FromStringOrNil("98c16e60-9b2a-45ff-a740-30243bc95ba6"),
		},
	},
}

func seedRoles(db *gorm.DB) error {
	if !db.HasTable(&types.Role{}) {
		if err := db.AutoMigrate(&types.Role{}).Error; err != nil {
			return err
		}
	}
	for _, role := range roles {
		if err := role.Create(db); err != nil {
			return err
		}
	}
	for roleID, permissions := range rolePermissions {
		var role = types.Role{
			ID: roleID,
		}
		if err := db.Model(role).Association("Permissions").Append(permissions).Error; err != nil {
			return err
		}
	}
	return nil
}
