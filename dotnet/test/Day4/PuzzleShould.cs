using System.IO;
using System.Linq;
using Shouldly;
using Xunit;

namespace AdventOfCode.Day4
{
    public class PuzzleShould
    {
        [Theory]
        [InlineData("aa bb cc dd ee", true)]
        [InlineData("aa bb cc dd aa", false)]
        [InlineData("aa bb cc dd aaa", true)]
        public void SolvePart1GivenExamples(string passphrase, bool expectedValidity)
        {
            // Act
            var actualValidity = Puzzle.SolvePart1(passphrase);

            // Assert
            actualValidity.ShouldBe(expectedValidity);
        }

        [Fact]
        public void SolvePart1()
        {
            // Arrange
            var passphrases = File.ReadAllLines(@"Day4\Input.txt");

            // Act
            var validPassphrases = passphrases
                .Where(Puzzle.SolvePart1)
                .ToArray();

            // Assert
            validPassphrases.Length.ShouldBe(451);
        }

        [Theory]
        [InlineData("abcde fghij", true)]
        [InlineData("abcde xyz ecdab", false)]
        [InlineData("a ab abc abd abf abj", true)]
        [InlineData("iiii oiii ooii oooi oooo", true)]
        [InlineData("oiii ioii iioi iiio", false)]
        public void SolvePart2GivenExamples(string passphrase, bool expectedValidity)
        {
            // Act
            var actualValidity = Puzzle.SolvePart2(passphrase);

            // Assert
            actualValidity.ShouldBe(expectedValidity);
        }

        [Fact]
        public void SolvePart2()
        {
            // Arrange
            var passphrases = File.ReadAllLines(@"Day4\Input.txt");

            // Act
            var validPassphrases = passphrases
                .Where(Puzzle.SolvePart2)
                .ToArray();

            // Assert
            validPassphrases.Length.ShouldBe(223);
        }
    }
}
