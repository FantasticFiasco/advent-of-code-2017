using System.IO;
using System.Linq;
using Shouldly;
using Xunit;

namespace AdventOfCode.Day5
{
    public class PuzzleShould
    {
        [Fact]
        public void SolvePart1GivenExample()
        {
            // Arrange
            var jumpInstructions = ToNumberArray("0\n3\n0\n1\n-3");
            
            // Act
            var actualSteps = Puzzle.SolvePart1(jumpInstructions);

            // Assert
            actualSteps.ShouldBe(5);
        }
        
        [Fact]
        public void SolvePart1()
        {
            // Arrange
            var jumpInstructions = ToNumberArray(File.ReadAllText(Path.Combine("Day5", "Input.txt")));
            
            // Act
            var actualSteps = Puzzle.SolvePart1(jumpInstructions);

            // Assert
            actualSteps.ShouldBe(388611);
        }
        
        [Fact]
        public void SolvePart2GivenExample()
        {
            // Arrange
            var jumpInstructions = ToNumberArray("0\n3\n0\n1\n-3");
            
            // Act
            var actualSteps = Puzzle.SolvePart2(jumpInstructions);

            // Assert
            actualSteps.ShouldBe(10);
        }
        
        [Fact]
        public void SolvePart2()
        {
            // Arrange
            var jumpInstructions = ToNumberArray(File.ReadAllText(Path.Combine("Day5", "Input.txt")));
            
            // Act
            var actualSteps = Puzzle.SolvePart2(jumpInstructions);

            // Assert
            actualSteps.ShouldBe(27763113);
        }
        
        private static int[] ToNumberArray(string jumpInstructions) =>
            jumpInstructions
                .Split("\n")
                .Select(int.Parse)
                .ToArray();
    }
}
