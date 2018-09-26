namespace AdventOfCode.Day1
{
    public static class PartTwo
    {
        public static int Solve(int[] input)
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
