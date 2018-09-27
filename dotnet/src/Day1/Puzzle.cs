namespace AdventOfCode.Day1
{
    public static class Puzzle
    {
        public static int SolvePart1(int[] input)
        {
            var sum = 0;

            for (var i = 0; i < input.Length; i++)
            {
                var current = input[i];
                var next = input[(i + 1) % input.Length];

                if (current == next)
                {
                    sum += current;
                }
            }

            return sum;
        }

        public static int SolvePart2(int[] input)
        {
            var sum = 0;

            for (var i = 0; i < input.Length; i++)
            {
                var current = input[i];
                var halfwayAround = input[(i + input.Length / 2) % input.Length];

                if (current == halfwayAround)
                {
                    sum += current;
                }
            }

            return sum;
        }
    }
}
