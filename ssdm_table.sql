CREATE TABLE SSDM_RAW(
   SSN varchar(9) PRIMARY KEY     NOT NULL,
   FIRSTNAME           CHAR(15)    NOT NULL,
   MIDDLENAME           CHAR(15)    NOT NULL,
   LASTNAME           CHAR(20)    NOT NULL,
   SUFFIX           varCHAR(6),
   BDAY            varchar(4),
   BMONTH            varchar(4),
   BYEAR            varchar(4),
   DDAY            varchar(4),
   DMONTH            varchar(4),
   DYEAR            varchar(4),
   RAW        varchar(82)
);

copy SSDM_RAW from '/home/ec2-user/output1.csv' DELIMITERS ',' CSV;