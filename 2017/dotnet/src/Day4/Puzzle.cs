using System.Linq;

namespace AdventOfCode.Day4
{
    public static class Puzzle
    {
        public static bool SolvePart1(string passphrase)
        {
            var words = passphrase.Split(' ');
            var distinctWords = words.Distinct().ToArray();

            return words.Length == distinctWords.Length;
        }

        public static bool SolvePart2(string passphrase)
        {
            var words = passphrase
                .Split(' ')
                .Select(word => word.ToCharArray())
                .Select(charArray => charArray.OrderBy(character => character))
                .Select(charArray => new string(charArray.ToArray()))
                .ToArray();

            var distinctWords = words.Distinct().ToArray();

            return words.Length == distinctWords.Length;
        }
    }
}
