//Mock testing for helpers
package database

// import (
// 	"context"
// 	"testing"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// //Mock collection
// type mockCollection struct {
// }

// //Mock collection implimenting dbIface.CollectionIface
// func (m *mockCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {

// 	c := &mongo.InsertOneResult{}
// 	return c, nil
// }

// func (m *mockCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (cur *mongo.Cursor, err error) {

// 	c := &mongo.Cursor{}
// 	return c, nil
// }

// func (m *mockCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {

// 	c := &mongo.UpdateResult{}
// 	return c, nil
// }

// func (m *mockCollection) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {

// 	c := &mongo.DeleteResult{}
// 	return c, nil
// }

// func (m *mockCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {

// 	c := &mongo.DeleteResult{}
// 	return c, nil
// }

// //Tests goes here...
// func TestCreateContact_Mock(t *testing.T) {

// 	_, err := CreateContact(contact, &mockCollection{})
// 	if err != nil {
// 		t.Error(`CreateContact Test FAIILED!
// 		err: `, err)
// 	}

// 	t.Log("CreateContact Test PASSED.")

// }

// func TestGetAllContacts_Mock(t *testing.T) {
// 	_, err := GetAllContacts(&mockCollection{})
// 	if err != nil {
// 		t.Error(`GetAllContacts Test FAILED
// 		err : `, err)
// 	}
// 	t.Logf("GetAllContacts Test PASSED.")

// }

// func TestUpdateContact_Mock(t *testing.T) {

// 	err := UpdateContact("627a5cf2d38f45fd608fdde7", contact.Name, contact.Number, &mockCollection{})
// 	if err != nil {
// 		t.Error(`UpdateContact Test FAILED!
// 		err: `, err)
// 	}
// 	t.Log("UpdateContact Test PASSED!")
// }

// func TestDeleteAllContacts_Mock(t *testing.T) {

// 	_, err := DeleteAllContacts(&mockCollection{})
// 	if err != nil {
// 		t.Error(`DeleteAllContacts Test FAILED!
// 		err: `, err)
// 	}

// 	t.Log("DeleteAllContacts Test PASSED.")

// }

// func TestDeleteContact_Mock(t *testing.T) {

// 	err := DeleteContact("627a5cf2d38f45fd608fdde7", &mockCollection{})
// 	if err != nil {
// 		t.Error(`DeleteContact Test FAILED!
// 		err: `, err)
// 	}

// 	t.Log("DeleteContact Test PASSED.")

// }
