# Gorbage

Gorbage is a simple garbage collector written in go.
The whole process is from the aritcle "https://journal.stuffwithstuff.com/2013/12/08/babys-first-garbage-collector/".
Its written in C. But I am re-implementing it in golang.

## What is GC

garbage collector is a simple program that provides you infinite memory. Just kidding. 
Its an illusion of infinite memory. 
What is does is after a certain threshold of you allocating memory for things in heap like 1000 arrays there is no more memory to allocate from heap.
So GC does it collect garbage.
Garbage in the sense the memory that is no longer in use.
So collect garbage -> free the space.
And there you have your memory.

"For the illusion of infinite memory to work, the language needs to be very safe about “no longer being used”. "


## Why we need GC

Well actually you don't need it. You yourself can allocate and free the memory space.
And also deal with all the problem that comes with poor memory management.
So to handle the memory allocation and deallocation we need the garbage collector.
Or else you have manually do it.
Great example for manual memory management is C. Where you use malloc and free.
Golang has garbage collector built into it.
And Rust uses a different method to handle memory management.

## When to worry about GC/memory allocation

Most of the time you don't need to give shit about GC or memory allocation.
Unless you see your application being in cloud nine.
Meaning there is a critical need for performance.

## Foundation of GC is Graph

So you have allocated some memory in heap. 
You have a pointers to those memory. These are the roots.
From that root you can connect to other memory spaces that take memory in heap.
Repeat that again and agian for every root and pointer you find you get a Graph.

## Algorithm for GC: Mark and Sweep

So we have our graph. Then we can run the algorithm on the graph
Simple steps really,
 - Starting a root mark the entire Graph. 
   Marking indicates that memory is in still use since you are referencing to it.
   Every time you reach an object, set a “mark” bit on it to true.
 - Once marking is done find all object that have not marked and delete them. 




 