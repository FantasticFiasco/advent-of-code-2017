using System.IO;
using System.Linq;
using Shouldly;
using Xunit;

namespace AdventOfCode.Day6
{
    public class PuzzleShould
    {
        [Fact]
        public void SolvePart1GivenExample()
        {
            // Act
            var redistributionCycles = Puzzle.SolvePart1(new[] { 0, 2, 7, 0 });

            redistributionCycles.ShouldBe(5);
        }

        [Fact]
        public void SolvePart1()
        {
            // Arrange
            var memoryBanks = ToNumberArray(File.ReadAllText(Path.Combine("Day6", "Input.txt")));

            // Act
            var redistributionCycles = Puzzle.SolvePart1(memoryBanks);

            redistributionCycles.ShouldBe(11137);
        }

        [Fact]
        public void SolvePart2GivenExample()
        {
            // Act
            var redistributionCycles = Puzzle.SolvePart2(new[] { 2, 4, 1, 2 });

            redistributionCycles.ShouldBe(4);
        }
        
        [Fact]
        public void SolvePart2()
        {
            // Act
            var redistributionCycles = Puzzle.SolvePart2(new[] { 14, 13, 12, 11, 9, 8, 8, 6, 6, 4, 4, 3, 1, 1, 0, 12 });

            redistributionCycles.ShouldBe(1037);
        }

        private static int[] ToNumberArray(string memoryBank) =>
            memoryBank
                .Split(",")
                .Select(int.Parse)
                .ToArray();
    }
}
