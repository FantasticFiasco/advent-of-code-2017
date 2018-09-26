namespace AdventOfCode.Day1
{
    public static class PartOne
    {
        public static int Solve(int[] input)
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
    }
}
