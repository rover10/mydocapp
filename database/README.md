
-- --------------------------------
-- Starts at 3:08 AM - 20 June 2020
-- --------------------------------

-- -----------
-- ENTITIES(onboarding)
-- -----------
-- User(id serial, uid uuid DEFAULT uuid_generate_v4(), user_type varchar, is_active boolean, first_name varchar, last_name varchar, gender varchar, phone varchar not null, email varchar not null, country_id integer references country(id), create_on datetime, updated_on datetime)
-- Patient(user_id integer references user(id), first_name varchar, last_name varchar, gender varchar, age int, existing_condition jsonb) -- existing_condition such as diabetes, heart issue, allergy
-- Doctor(user_id integer references user(id), fee number, rating number, practice_start_date datetime, approved boolean)
-- Clinic( uid uuid DEFAULT uuid_generate_v4(), name varchar, address varchar, state_id integer refernces states(id), phone varchar not null, email varchar not null)
-- Staff(user_id integer references user(id), clinic_id uuid references clinic(uid))


-- --------------------------
-- MASTER
-- --------------------------
-- Disease(id serial, name varchar unique, description varchar)
-- Doctor_category(id serial, name varchar unique not null) (mbbs, md, optothamly, surgen, consultant)
-- Clinic_type(id serial, type varchar unique not null) (gov,private,charitable,ngo,corporate)
-- Qualification(id serial, qualification varchar unique not null)
-- Tests(id serial, type varchar unique not null, pre_requisite varchar, description varchar, price number)
-- Role(id serial, type varchar unique not null
-- UserType(id serial, type varchar unique not null)
-- Document_type(id serial, type varchar unique not null) (dr_degree, dr_licence, voter_id, passport, patient_report)


-- ----------------------------
-- Operations
-- ----------------------------
-- Booking(uid uuid DEFAULT uuid_generate_v4(), user_id integer references user(id), doctor_id uuid references user(uid), disease_id integer references disease(id), no_show boolean, booking_date datetime, slot_date_time datetime, contact_phone varchar) -- no_show patient reported
-- TreatmentDetail(booking_id uuid references booking(uid), doctor_id uuid references user(uid), date_created datetime, patient_problem_description jsonb, doctor_diganosis jsonb, prescription jsonb, test jsonb) -- Multiple doctor can treat a paitent that is why we store doctor id
-- DoctorReviews (booking_id uuid references booking(uid), doctor_id integer references user(id),reviewer_id integer references user(id), rating number, review varchar, review_date datetime)
-- TestReport(uid uuid DEFAULT uuid_generate_v4(), booking_id uuid references booking(uid), created_on datetime, updated_on date_time) 
-- StaffRole(user_id integer references user(id), role_id integer references role(id))
-- DoctorQualification(user_uid uuid references user(uid) not null, qualification_id integer references qualification(id))
-- UserDocuments(user_uid uuid references user(uid) not null, doc_type_id integer not null, url varchar not null)


-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- uid uuid DEFAULT uuid_generate_v4()

-- --------------------------------
-- Starts at 5:02 AM - 20 June 2020
-- --------------------------------

