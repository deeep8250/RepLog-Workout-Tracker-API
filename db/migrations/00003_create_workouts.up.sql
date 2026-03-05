create table if not exists workouts(
    id serial primary key,
    user_id int references users(id) on delete cascade,
    notes text,
    date timestamp default now()
);

create index if not exists idx_workouts_user_id on workouts(user_id);
