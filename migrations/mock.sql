INSERT INTO users (first_name, last_name, avatar_link, gender, age, phone_number, city_of_residence, description, email, is_email_verified, username, is_private, receive_push,role)
VALUES
('Rakhat','Oshakbayev','link',true,22,'87756400316','Astana','description','rakhat.oshakbayev@gmail.com',true,'roshakba',true,true,'admin')
('John', 'Doe', 'avatar1.jpg', true, 30, '123456789', 'New York', 'Description 1', 'john.doe@example.com', true, 'johndoe', false, true,'user'),
('Jane', 'Smith', 'avatar2.jpg', false, 25, '987654321', 'Los Angeles', 'Description 2', 'jane.smith@example.com', true, 'janesmith', false, true,'user');

INSERT INTO blocks (who_blocked_id,block_id) 
VALUES
(1,2),
(3,1);

INSERT INTO friend_requests (from_user_id,to_user_id)
VALUES
(1,2,'pending', CURRENT_TIMESTAMP,CURRENT_TIMESTAMP),
(2,3,'rejected', CURRENT_TIMESTAMP,CURRENT_TIMESTAMP);

INSERT INTO friendships (user_id1,user_id2)
VALUES ()