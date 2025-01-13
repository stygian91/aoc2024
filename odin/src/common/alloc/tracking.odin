package tracking

import "base:runtime"
import "core:fmt"
import "core:mem"

print_leaks :: proc(alloc: ^mem.Tracking_Allocator) {
	fmt.eprintfln("------------------")
	fmt.eprintfln("Leaks:")

	for _, entry in alloc.allocation_map {
		fmt.eprintfln("- %v leaked %v bytes", entry.location, entry.size)
	}

	fmt.eprintfln("Bad frees:")

	for entry in alloc.bad_free_array {
		fmt.eprintfln("- %v bad free", entry.location)
	}
}

track_leaks_for_proc :: proc(callback: proc()) {
	track_alloc: mem.Tracking_Allocator
	mem.tracking_allocator_init(&track_alloc, context.allocator)
	context.allocator = mem.tracking_allocator(&track_alloc)

	defer {
		print_leaks(&track_alloc)
		mem.tracking_allocator_destroy(&track_alloc)
	}

  callback()
}
