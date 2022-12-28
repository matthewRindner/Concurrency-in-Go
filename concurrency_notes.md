Concurreny != Parallelism

Parallelism:
    - to run two thing at exactly the same time
    - on a multi-core system, core_1 would run on task while core_2 would do another task at the same time

- Issue is that typically lines of code must be run in the right order, so it's hard to execute 2 lines at the same time, or parellized. 

Concurreny:
    - Breaking up a program into independently executable tasks that could be run at the same time and yeild the correct result at the end
    -a concurrent program CAN be parallelized
    - This is the strong seeling point of GO ny focussing on the struclure of the code without worrying about what is running on which core
    

