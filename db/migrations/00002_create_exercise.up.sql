create table if not exists exercises(

id serial primary key,
name varchar(255) not null,
muscle_group varchar(255) not null,
created_by int references users(id) on delete cascade,   
created_at timestamp default now()
);

create index if not exists idx_exercise_muscle_group on exercises(muscle_group);
create index if not exists idx_exercise_created_by on exercises(created_by);
