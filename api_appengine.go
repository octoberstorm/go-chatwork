package gochatwork

import (
	"encoding/json"
)

func (c *AppengineClient) Me() Me {
	ret := c.Get("/me", map[string]string{})
	var me Me
	json.Unmarshal(ret, &me)
	return me
}

// params keys
//  - assigned_by_account_id
//  - status: [open, done]
func (c *AppengineClient) MyTasks(params map[string]string) []MyTask {
	ret := c.Get("/my/tasks", params)
	var tasks []MyTask
	json.Unmarshal(ret, &tasks)
	return tasks
}

func (c *AppengineClient) Contacts() []Contact {
	ret := c.Get("/contacts", map[string]string{})
	var contacts []Contact
	json.Unmarshal(ret, &contacts)
	return contacts
}

func (c *AppengineClient) Rooms() []Room {
	ret := c.Get("/rooms", map[string]string{})
	var rooms []Room
	json.Unmarshal(ret, &rooms)
	return rooms
}

func (c *AppengineClient) Room(roomId string) Room {
	ret := c.Get("/rooms/"+roomId, map[string]string{})
	var room Room
	json.Unmarshal(ret, &room)
	return room
}

// params keys
//   * name
//   * members_admin_ids
//   - description
//   - icon_preset
//   - members_member_ids
//   - members_readonly_ids
func (c *AppengineClient) CreateRoom(params map[string]string) []byte {
	return c.Post("/rooms", params)
}

// params keys
//   - description
//   - icon_preset
//   - name
func (c *AppengineClient) UpdateRoom(roomId string, params map[string]string) []byte {
	return c.Put("/rooms/"+roomId, params)
}

// params key
//   * action_type: [leave, delete]
func (c *AppengineClient) DeleteRoom(roomId string, params map[string]string) []byte {
	return c.Delete("/rooms/"+roomId, params)
}

func (c *AppengineClient) RoomMembers(roomId string) []Member {
	ret := c.Get("/rooms/"+roomId+"/members", map[string]string{})
	var members []Member
	json.Unmarshal(ret, &members)
	return members
}

// params keys
//   * members_admin_ids
//   - members_member_ids
//   - members_readonly_ids
func (c *AppengineClient) UpdateRoomMembers(roomId string, params map[string]string) []byte {
	return c.Put("/rooms/"+roomId+"/members", params)
}

func (c *AppengineClient) RoomMessages(roomId string) []Message {
	ret := c.Get("/rooms/"+roomId+"/messages", map[string]string{})
	var messages []Message
	json.Unmarshal(ret, &messages)
	return messages
}

func (c *AppengineClient) PostRoomMessage(roomId string, body string) []byte {
	return c.Post("/rooms/"+roomId+"/messages", map[string]string{"body": body})
}

func (c *AppengineClient) RoomMessage(roomId, messageId string) Message {
	ret := c.Get("/rooms/"+roomId+"/messages/"+messageId, map[string]string{})
	var message Message
	json.Unmarshal(ret, &message)
	return message
}

func (c *AppengineClient) RoomTasks(roomId string) []Task {
	ret := c.Get("/rooms/"+roomId+"/tasks", map[string]string{})
	var tasks []Task
	json.Unmarshal(ret, &tasks)
	return tasks
}

// params keys
//   * body
//   * to_ids
//   - limit
func (c *AppengineClient) PostRoomTask(roomId string, params map[string]string) []byte {
	return c.Post("/rooms/"+roomId+"/tasks", params)
}

func (c *AppengineClient) RoomTask(roomId, taskId string) Task {
	ret := c.Get("/rooms/"+roomId+"/tasks/"+taskId, map[string]string{})
	var task Task
	json.Unmarshal(ret, &task)
	return task
}

// params key
//   - account_id
func (c *AppengineClient) RoomFiles(roomId string, params map[string]string) []File {
	ret := c.Get("/rooms/"+roomId+"/files", params)
	var files []File
	json.Unmarshal(ret, &files)
	return files
}

func (c *AppengineClient) RoomFile(roomId, fileId string) File {
	ret := c.Get("/rooms/"+roomId+"/files/"+fileId, map[string]string{})
	var file File
	json.Unmarshal(ret, &file)
	return file
}
