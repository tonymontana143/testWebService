begin; 
create table employees (
    id uuid primary key default gen_random_uuid(),
    full_name varchar(100) not null,
    phone_number varchar(15) not null,
    city varchar(50) not null
);
commit;