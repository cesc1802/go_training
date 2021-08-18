### Yêu cầu chức năng:
- User có thể tạo mới task. Mỗi ngày user được phép tạo tối đa N task (ví dụ: 1 ngày tạo tối đa 10 task)
- Nếu user đã tạo max số task được tạo trong ngày. trả về lỗi 4xx cho client và từ chối request tạo mới task đó.
### Yêu cầu phi chức năng:
- Fork repository, và sử dụng PR khi bạn phát triển xong.
- Sử dụng một loại db khác mà bạn thành thao nhất thay thế sqlite (MySql, Postgree, Oracle).
- Phân chia thành các layer rõ ràng (service layer, use-case layer, storage layer)


#### DB Schema
```sql
-- users definition

CREATE TABLE users (
	id TEXT NOT NULL,
	password TEXT NOT NULL,
	max_todo INTEGER DEFAULT 5 NOT NULL,
	CONSTRAINT users_PK PRIMARY KEY (id)
);

INSERT INTO users (id, password, max_todo) VALUES('firstUser', 'example', 5);

-- tasks definition

CREATE TABLE tasks (
	id TEXT NOT NULL,
	content TEXT NOT NULL,
	user_id TEXT NOT NULL,
    created_date TEXT NOT NULL,
	CONSTRAINT tasks_PK PRIMARY KEY (id),
	CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
);
```

#### Sequence diagram
![auth and create tasks request](https://github.com/cesc1802/go_training/blob/master/docs/sequence.svg)
