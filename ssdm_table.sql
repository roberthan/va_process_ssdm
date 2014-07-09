CREATE TABLE SSDM_RAW(
   SSN INT PRIMARY KEY     NOT NULL,
   FIRSTNAME           CHAR(15)    NOT NULL,
   MIDDLENAME           CHAR(15)    NOT NULL,
   LASTNAME           CHAR(20)    NOT NULL,
   SUFFIX           CHAR(4),
   BDAY            SMALLINT,
   BMONTH            SMALLINT,
   BYEAR            SMALLINT,
   DDAY            SMALLINT,
   DMONTH            SMALLINT,
   DYEAR            SMALLINT,
   RAW        CHAR(82) NOT NULL
);