create table disease(id serial primary key, name varchar unique, description varchar);
create table doctor_category(id serial primary key, name varchar unique not null, description varchar); -- (mbbs, md, optothamly, surgen, consultant)
create table clinic_type(id serial primary key, type varchar unique not null); -- (gov,private,charitable,ngo,corporate)
create table qualification(id serial primary key, qualification varchar unique not null); -- (mbbs, certification)
create table test(id serial primary key, type varchar unique not null, prerequisite varchar, description varchar, price float);
create table roles(id serial primary key, type varchar unique not null); -- staff role such as admin assistant, pathologist
create table user_type(id serial primary key, type varchar unique not null); -- doctor, patient, staff
create table document_type(id serial primary key, type varchar unique not null); -- (dr_degree, dr_licence, voter_id, passport, patient_report)
create table gender(id serial primary key, type varchar unique not null); -- male, female, lgbt
create table country(id serial primary key, name varchar not null, code integer not null, abbreviation varchar not null);
create table state(id serial primary key, name varchar not null, country_id integer references country(id));
