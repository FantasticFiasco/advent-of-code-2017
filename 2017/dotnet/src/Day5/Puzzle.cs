namespace AdventOfCode.Day5
{
    public static class Puzzle
    {
        public static int SolvePart1(int[] jumpInstructions)
        {
            int currentIndex = 0;
            int steps = 0;
            
            while (currentIndex > -1 && currentIndex < jumpInstructions.Length)
            {
                var nextIndex = currentIndex + jumpInstructions[currentIndex]; 
                
                // Increment current instruction by 1
                jumpInstructions[currentIndex]++;

                currentIndex = nextIndex;
                steps++;
            }

            return steps;
        }
        
        public static int SolvePart2(int[] jumpInstructions)
        {
            int currentIndex = 0;
            int steps = 0;
            
            while (currentIndex > -1 && currentIndex < jumpInstructions.Length)
            {
                var nextIndex = currentIndex + jumpInstructions[currentIndex];

                // Decrement or increment current instruction depending on its value
                if (jumpInstructions[currentIndex] >= 3)
                {
                    jumpInstructions[currentIndex]--;
                }
                else
                {
                    jumpInstructions[currentIndex]++;    
                }

                currentIndex = nextIndex;
                steps++;
            }

            return steps;
        }
    }
}
