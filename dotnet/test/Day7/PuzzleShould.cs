using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;
using Shouldly;
using Xunit;

namespace AdventOfCode.Day7
{
    public class PuzzleShould
    {
        private static readonly Regex InputRegex = new Regex(
            @"^(?<name>\w+) \((?<weight>\d+)\)( -> (?<programsAbove>.*))?$",
            RegexOptions.Singleline);
        
        [Fact]
        public void SolvePart1GivenExample()
        {
            // Arrange
            var programs = new[]
            {
                new Program("pbga", 66),
                new Program("xhth", 57),
                new Program("ebii", 61),
                new Program("havc", 66),
                new Program("ktlj", 57),
                new Program("fwft", 72, "ktlj", "cntj", "xhth"),
                new Program("qoyq", 66),
                new Program("padx", 45, "pbga", "havc", "qoyq"),
                new Program("tknk", 41, "ugml", "padx", "fwft"),
                new Program("jptl", 61),
                new Program("ugml", 68, "gyxo", "ebii", "jptl"),
                new Program("gyxo", 61),
                new Program("cntj", 57)
            };

            // Act
            var bottomProgram = Puzzle.SolvePart1(programs);
            
            // Assert
            bottomProgram.Name.ShouldBe("tknk");
        }

        [Fact]
        public void SolvePart1()
        {
            // Arrange
            var programs = ToPrograms(File.ReadAllLines(Path.Combine("Day7", "Input.txt")));
            
            // Act
            var bottomProgram = Puzzle.SolvePart1(programs);
            
            // Assert
            bottomProgram.Name.ShouldBe("hmvwl");
        }

        private static Program[] ToPrograms(IEnumerable<string> rows) =>
            rows
                .Select(row => InputRegex.Match(row))
                .Select(match =>
                {
                    var name = match.Groups["name"].Value;
                    var weight = int.Parse(match.Groups["weight"].Value);
                    var programsAbove = match.Groups["programsAbove"].Success
                        ? match.Groups["programsAbove"].Value.Split(",").Select(programAbove => programAbove.Trim())
                        : Enumerable.Empty<string>();
                    
                    return new Program(name, weight, programsAbove.ToArray());
                })
                .ToArray();
    }
}
