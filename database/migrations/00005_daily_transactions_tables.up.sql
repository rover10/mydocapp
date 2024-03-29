create table appointment(uid uuid DEFAULT uuid_generate_v4() primary key, clinic_id uuid references clinic(uid) not null, account_id uuid references users(uid) not null, patient_id uuid references patient(uid) not null, doctor_id uuid references users(uid), [disease_id integer references disease(id)], no_show boolean default false not null, created_on timestamp default now(), updated_on timestamp, slot_date_time timestamp not null, contact_phone varchar); -- no_show patient reported
create table treatment(uid uuid DEFAULT uuid_generate_v4() primary key, appointment_id uuid references appointment(uid) not null, doctor_id uuid references users(uid) not null, patient_problem_description varchar, created_on timestamp default now());
create table doctor_observation(treatment_id uuid references treatment(uid) not null, observation varchar, created_on timestamp default now() not null, doctor_id uuid references users(uid) not null);
create table prescription(treatment_id uuid references treatment(uid) not null, prescription varchar, created_on timestamp default now() not null, doctor_id uuid references users(uid) not null, note varchar not null, is_active boolean);
create table patient_test(treatment_id uuid references treatment(uid) not null, test int references test(id) not null, description varchar, doc_url varchar, created_on timestamp not null default now(), doctor_id uuid references users(uid) not null);

create table doctor_review(appointment_id uuid references appointment(uid), reviewer_id uuid references users(uid), doctor_id uuid references users(uid), rating float check (rating >= 0) check (rating <= 5), review varchar, created_on timestamp default now());
create unique index if not exists unique_appointment_id_and_doctor_id on doctor_review(appointment_id, doctor_id);

create table user_document(uid uuid DEFAULT uuid_generate_v4() primary key, user_id uuid references users(uid) not null, doc_type_id integer not null, url varchar not null, created_on timestamp default now() not null);
create table test_report(uid uuid DEFAULT uuid_generate_v4() primary key, appointment_id uuid references appointment(uid), created_on timestamp, updated_on timestamp, doc_id uuid references user_document(uid)); 
create table staff_role(user_id uuid references users(uid) not null,clinic_id uuid references clinic(uid) not null, role_id integer references roles(id) not null, created_on timestamp default now() not null, is_active boolean not null default true);
-- one staff has disticnt role
create unique index if not exists unique_staff_id_role_id_staff_role on staff_role(user_id, role_id);

create table doctor_qualification(user_id uuid references users(uid) not null, qualification_id integer references qualification(id) not null, created_on timestamp default now() not null, certificate_doc uuid references user_document(uid) not null, verified boolean not null default false);
  
drop table appointment , treatment_detail , doctor_review , users, test_report , staff_role , doctor_qualification cascade;