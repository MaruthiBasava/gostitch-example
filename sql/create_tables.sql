

create table dog_intelligence(
    dog_intel_id uuid default generate_uuid_v4(),
    breed varchar(50) not null,
    classification varchar(100) not null,
    obey float,
    reps_lower int,
    reps_higher int,
    primary key(dog_intel_id),
    unique(breed)
);

create table dog_breeds(
    dog_breed_id uuid default generate_uuid_v4(),
    breed varchar(50) not null,
    height_low_inches int, 
    height_high_inches int,
    weight_low_lbs int,
    weight_high_lbs int,
    primary key (dog_breed_id),
    unique(breed)
);


