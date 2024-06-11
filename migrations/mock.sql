INSERT INTO users (first_name, last_name, password, avatar_link, gender, age, phone_number, city_of_residence, description, email, is_email_verified, username, is_private, receive_push,role_id,created_at)
VALUES
('Rakhat','Oshakbayev','$2a$10$Fvk3WTJiY3rFzEigO6QemOEgSaIOK9Hup.ExkR0HBl6BK67xQmBEe','link',true,22,'87756400316','Astana','description','rakhat.oshakbayev@gmail.com',true,'roshakba',true,true,1,CURRENT_TIMESTAMP),
('John', 'Doe', '$2a$10$Fvk3WTJiY3rFzEigO6QemOEgSaIOK9Hup.ExkR0HBl6BK67xQmBEe','avatar1.jpg', true, 30, '123456789', 'New York', 'Description 1', 'john.doe@example.com', true, 'johndoe', false, true,2,CURRENT_TIMESTAMP),
('Jane', 'Smith', '$2a$10$Fvk3WTJiY3rFzEigO6QemOEgSaIOK9Hup.ExkR0HBl6BK67xQmBEe','avatar2.jpg', false, 25, '987654321', 'Los Angeles', 'Description 2', 'jane.smith@example.com', true, 'janesmith', false, true,2,CURRENT_TIMESTAMP);

INSERT INTO blocks (who_blocked_id,blocked_id) 
VALUES
(1,2),
(3,1);

INSERT INTO friend_requests (requestor_id,recipient_id,status,created_at)
VALUES
(1,2,'pending', CURRENT_TIMESTAMP),
(2,3,'rejected', CURRENT_TIMESTAMP);

INSERT INTO friendships (user_id1,user_id2)
VALUES (1,3);