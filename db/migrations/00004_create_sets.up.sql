create table if not exists sets(
    id serial primary key,
    workout_id int references workouts(id) on delete cascade,
    exercise_id int references exercises(id) on delete cascade,
    reps int not null,
    weight_kg int not null
);

create index if not exists idx_sets_workout_id on sets(workout_id);
create index if not exists idx_sets_exercise_id on sets(exercise_id);