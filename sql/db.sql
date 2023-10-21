create table customer (
    id              SERIAL PRIMARY KEY,
    firstname       text NOT NULL,
	lastname        text NOT NULL,
    mobile_no       text NOT NULL,
    the1_member_id  text NULL,
	the1_mobile_no  text NULL,
    is_active       bool NOT NULL DEFAULT true,
    created_at      timestamptz NOT NULL,
	created_by      text NOT NULL,
	updated_at      timestamptz NOT NULL,
	updated_by      text NOT NULL
);

insert into customer (firstname, lastname, mobile_no, created_at, created_by, updated_at, updated_by)
values('test_name_01', 'test_lastname_01', '0912345678', now(), 'admin', now(), 'admin');

create table address (
    id          SERIAL PRIMARY KEY,
    "owner"     integer NOT NULL,
    subdistrict text  NOT NULL,
    district text  NOT NULL,
    province    text NOT NULL,
    CONSTRAINT fk_customer
      FOREIGN KEY("owner") 
	  REFERENCES customer(id)
);