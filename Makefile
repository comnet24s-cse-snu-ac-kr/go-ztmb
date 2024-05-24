CC = g++
CXXFLAGS = -Wall -O2
OBJS = logic.o

logic : $(OBJS)
	$(CC) $(CXXFLAGS) $(OBJS) -o logic

logic.o : logic.cc
	$(CC) $(CXXFLAGS) -c logic.cc
