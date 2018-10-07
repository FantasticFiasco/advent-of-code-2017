using System.Collections.Generic;
using System.Linq;

namespace AdventOfCode.Day7
{
    public static class Puzzle
    {
        public static Program SolvePart1(Program[] programs)
        {
            return programs.Single(program => !IsProgramAboveOtherProgram(program, programs));
        }

        private static bool IsProgramAboveOtherProgram(Program program, IEnumerable<Program> programs) =>
            programs.Any(p => p.ProgramsAbove.Contains(program.Name));
    }
}
