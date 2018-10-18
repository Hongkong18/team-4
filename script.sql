CREATE DATABASE mindhk;
use mindhk;

create TABLE institutions
(
     id INT unsigned NOT NULL AUTO_INCREMENT, # Unique ID for the record
     name   VARCHAR(150) NOT NULL,                # Name of the cat
     targeting_group   VARCHAR(150) NOT NULL,  #facing groups
     location   VARCHAR(150) NOT NULL,                # Localtion of the institution
     description    VARCHAR(150) NOT NULL,                # description
     contact_info   VARCHAR(150) NOT NULL,     # Contact number
     status         VARCHAR(150) NOT NULL,      # The status of the company
     PRIMARY    KEY (id)

);

show tables;

INSERT INTO institutions ( name, targeting_group,location, description, contact_info,status) VALUES
  ( 'Alcoholics Anonymous Hong Kong','Youth', 'Hong Kong Island', 'Drinking problems', '+852 9073 6922','Approved' ),
  ( 'Alliance of Ex-mentally ill Hong Kong','Adults', 'Kowloon', 'ex-mentally ill', '+852 4311 9538' ,'Approved'),
  ( 'Early psychosis Foundation', 'Elders','', 'pyschosis','+852 6075 6504','Approved' ),
  ('Fu Hong Society', 'Adults, Youth','Hong Kong Island', 'Befriend people with disablity', '+852 9073 1234', 'Pending'),
  ( 'Kely Support','Youth', 'Hong Kong Island', 'Drug abuses', '+852 9022 1213', 'Pending'),
  ( 'The Samaritans', 'Adult','Kowloon', 'Social ability','+852 6075 6504' ,'Not Approved');
    

ALTER TABLE institutions ADD opening_hours VARCHAR(150) AFTER location;
update institutions
set opening_hours='9:30-19:30'
where id = 1;
update institutions
set opening_hours='9:30-19:30'
where id = 2;
update institutions
set opening_hours='9:30-19:30'
where id = 3;
update institutions
set opening_hours='9:30-19:30'
where id = 4;
update institutions
set opening_hours='9:30-19:30'
where id = 5;

ALTER TABLE institutions ADD languages VARCHAR(150) AFTER location;
update institutions
set languages='English/Mandarin/Cantonese'
where id = 1;

update institutions
set languages='English/Mandarin'
where id = 2;
update institutions
set languages='English/Cantonese'
where id = 3;
update institutions
set languages='English/Cantonese'
where id = 4;
update institutions
set languages='Mandarin/Cantonese'
where id = 5;


