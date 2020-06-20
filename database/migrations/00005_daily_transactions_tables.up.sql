-- ----------------------------
-- Operations
-- ----------------------------
	create table booking(uid uuid DEFAULT uuid_generate_v4() primary key, account_id uuid references users(uid) not null, doctor_id uuid references users(uid), disease_id integer references disease(id), no_show boolean default false, booking_date timestamp, slot_date_time timestamp, contact_phone varchar); -- no_show patient reported
	create table treatment_detail(booking_id uuid references booking(uid), doctor_id uuid references users(uid), date_created timestamp, patient_problem_description jsonb, doctor_diganosis jsonb, prescription jsonb, test jsonb); -- Multiple doctor can treat a paitent that is why we store doctor id
 	create table doctor_review(booking_id uuid references booking(uid), reviewer_id uuid references users(uid), doctor_id uuid references users(uid), rating float check (rating >= 0) check (rating <= 5), review varchar, review_date timestamp); -- unique(booking_id, doctor_id) single_reiew per booking
 	create unique index unique_booking_id_and_doctor_id on doctor_review(booking_id, doctor_id);
 
  	create table user_document(uid uuid DEFAULT uuid_generate_v4() primary key, user_uid uuid references users(uid) not null, doc_type_id integer not null, url varchar not null, date_created timestamp);
 	create table test_report(uid uuid DEFAULT uuid_generate_v4() primary key, booking_id uuid references booking(uid), created_on timestamp, updated_on timestamp, doc_id uuid references user_document(uid)); 
	create table staff_role(user_id uuid references users(uid) not null, role_id integer references roles(id) not null, date_created timestamp default now() not null, is_active boolean not null default true);
 	create table doctor_qualification(user_uid uuid references users(uid) not null, qualification_id integer references qualification(id), date_created timestamp default now() not null, certificate_doc uuid references user_document(uid), verified boolean not null default false);
 
  
