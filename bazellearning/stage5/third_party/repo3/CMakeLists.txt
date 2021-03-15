cmake_minimum_required(VERSION 3.5)
project(hello)
add_library( HelloUniverse SHARED helloUniverse.cc)
add_library( hello::library ALIAS HelloUniverse)

target_include_directories(HelloUniverse 
    PUBLIC
        ${PROJECT_SOURCE_DIR}/include
)
add_executable(hello main.cc)

target_link_libraries(${PROJECT_NAME}
    PRIVATE
        hello::library
)

if( CMAKE_INSTALL_PREFIX_INITIALIZED_TO_DEFAULT )
    message(STATUS "Set ...")
    set(CMAKE_INSTALL_PREFIX "${CMAKE_BINARY_DIR}/install" CACHE STRING "the path to use for make install" FORCE)
endif()
install(TARGETS hello DESTINATION bin )
install(TARGETS HelloUniverse  LIBRARY DESTINATION lib)
install(DIRECTORY ${PROJECT_SOURCE_DIR}/include DESTINATION include )