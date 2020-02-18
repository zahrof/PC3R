################################################################################
# Automatically-generated file. Do not edit!
################################################################################

# Add inputs and outputs from these tool invocations to the build variables 
C_SRCS += \
../src/Tapis.c \
../src/Threads_structures_and_routines.c \
../src/auxiliar_methods.c \
../src/paquet.c \
../src/tme2.c 

OBJS += \
./src/Tapis.o \
./src/Threads_structures_and_routines.o \
./src/auxiliar_methods.o \
./src/paquet.o \
./src/tme2.o 

C_DEPS += \
./src/Tapis.d \
./src/Threads_structures_and_routines.d \
./src/auxiliar_methods.d \
./src/paquet.d \
./src/tme2.d 


# Each subdirectory must supply rules for building sources it contributes
src/%.o: ../src/%.c
	@echo 'Building file: $<'
	@echo 'Invoking: GCC C Compiler'
	gcc -I/home/zahrof/Downloads/ft_v1.0_src/ft_v1.0/include -include/home/zahrof/Downloads/ft_v1.0_src/ft_v1.0/include/fthread.h -O0 -g3 -Wall -c -fmessage-length=0 -MMD -MP -MF"$(@:%.o=%.d)" -MT"$(@)" -o "$@" "$<"
	@echo 'Finished building: $<'
	@echo ' '


