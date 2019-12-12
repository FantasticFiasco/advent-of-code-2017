using System;
using System.Collections.Generic;
using System.Linq;

namespace AdventOfCode.Day6
{
    public static class Puzzle
    {   
        public static int SolvePart1(int[] memoryBanks)
        {
            var redistributionCycles = 0;

            var history = new History();

            while (!history.Contains(memoryBanks))
            {
                history.Add(memoryBanks);

                var index = Array.IndexOf(memoryBanks, memoryBanks.Max());

                Reallocate(memoryBanks, index);

                redistributionCycles++;
            }
            
            return redistributionCycles;
        }

        public static int SolvePart2(int[] memoryBanks)
        {
            // The solution to part 2 is the same as to part 1, the only thing that differs is the
            // initial memory banks
            return SolvePart1(memoryBanks);
        }
        
        private static void Reallocate(int[] memoryBanks, int index)
        {
            var blockCount = memoryBanks[index];

            // Clear block count
            memoryBanks[index] = 0;

            // Reallocate
            for (int offset = 1; offset <= blockCount; offset++)
            {
                var currentIndex = (index + offset) % memoryBanks.Length;

                memoryBanks[currentIndex]++;
            }
        }

        private class History
        {
            private readonly List<string> history;

            public History()
            {
                history = new List<string>();
            }

            public void Add(int[] memoryBanks) =>
                history.Add(ToString(memoryBanks));

            public bool Contains(int[] memoryBanks) =>
                history.Contains(ToString(memoryBanks));

            private static string ToString(int[] memoryBanks) =>
                string.Join(",", memoryBanks);
        }
    }
}
