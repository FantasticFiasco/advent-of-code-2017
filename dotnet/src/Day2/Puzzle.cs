using System.Linq;

namespace AdventOfCode.Day2
{
    public static class Puzzle
    {
        public static int SolvePart1(int[][] input)
        {
            var checksum = 0;

            foreach (var rowOfNumbers in input)
            {
                var min = rowOfNumbers.Min();
                var max = rowOfNumbers.Max();

                checksum += max - min;
            }

            return checksum;
        }

        public static int SolvePart2(int[][] input)
        {
            var checksum = 0;

            foreach (var rowOfNumbers in input)
            {
                for (int i = 0; i < rowOfNumbers.Length; i++)
                for (int j = 0; j < rowOfNumbers.Length; j++)
                {
                    if (i != j &&
                        rowOfNumbers[j] != 0 &&
                        rowOfNumbers[i] % rowOfNumbers[j] == 0)
                    {
                        checksum += rowOfNumbers[i] / rowOfNumbers[j];
                    }
                }
            }

            return checksum;
        }
    }
}
