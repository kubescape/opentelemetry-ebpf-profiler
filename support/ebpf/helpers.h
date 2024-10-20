// Provide helpers for the eBPF code.

#ifndef OPTI_HELPERS_H
#define OPTI_HELPERS_H

typedef enum ProgramType {
    PerfEvent,
    Kprobe,
} ProgramType;

// Macros for BPF program type and context handling.
#define DEFINE_DUAL_PROGRAM(name, func)                    \
SEC("perf_event/" #name)                                         \
int name##_perf(struct pt_regs *ctx)                         \
{                                                                        \
    return func(ctx, PerfEvent);                            \
}                                                                        \
                                                                         \
SEC("kprobe/" #name)                                             \
int name##_kprobe(struct pt_regs *ctx)                                   \
{                                                                        \
    return func(ctx, Kprobe);                                                    \
}


// get_unwinder_by_program_type returns the unwinder program index for the given program type.
static inline __attribute__((__always_inline__))
TracePrograms get_unwinder_by_program_type(ProgramType programType, TracePrograms unwinder) {
    switch (programType) {
        case Kprobe:
            switch (unwinder) {
                case PROG_UNWIND_STOP:
                    return PROG_KPROBE_UNWIND_STOP;
                case PROG_UNWIND_NATIVE:
                    return PROG_KPROBE_UNWIND_NATIVE;
                case PROG_UNWIND_HOTSPOT:
                    return PROG_KPROBE_UNWIND_HOTSPOT;
                case PROG_UNWIND_PERL:
                    return PROG_KPROBE_UNWIND_PERL;
                case PROG_UNWIND_PYTHON:
                    return PROG_KPROBE_UNWIND_PYTHON;
                case PROG_UNWIND_PHP:
                    return PROG_KPROBE_UNWIND_PHP;
                case PROG_UNWIND_RUBY:
                    return PROG_KPROBE_UNWIND_RUBY;
                case PROG_UNWIND_V8:
                    return PROG_KPROBE_UNWIND_V8;
                case PROG_UNWIND_DOTNET:
                    return PROG_KPROBE_UNWIND_DOTNET;
                default:
                    return -1;
            }
        case PerfEvent:
            switch (unwinder) {
                case PROG_UNWIND_STOP:
                    return PROG_UNWIND_STOP;
                case PROG_UNWIND_NATIVE:
                    return PROG_UNWIND_NATIVE;
                case PROG_UNWIND_HOTSPOT:
                    return PROG_UNWIND_HOTSPOT;
                case PROG_UNWIND_PERL:
                    return PROG_UNWIND_PERL;
                case PROG_UNWIND_PYTHON:
                    return PROG_UNWIND_PYTHON;
                case PROG_UNWIND_PHP:
                    return PROG_UNWIND_PHP;
                case PROG_UNWIND_RUBY:
                    return PROG_UNWIND_RUBY;
                case PROG_UNWIND_V8:
                    return PROG_UNWIND_V8;
                case PROG_UNWIND_DOTNET:
                    return PROG_UNWIND_DOTNET;
                default:
                    return -1;
            }
        default:
            return -1;
    }
}

#endif // OPTI_HELPERS_H