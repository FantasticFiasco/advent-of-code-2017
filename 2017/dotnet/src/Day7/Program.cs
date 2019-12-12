namespace AdventOfCode.Day7
{
    public class Program
    {
        public Program(string name, int weight, params string[] programsAbove)
        {
            Name = name;
            Weight = weight;
            ProgramsAbove = programsAbove;
        }
        
        public string Name { get; }

        public int Weight { get; }

        public string[] ProgramsAbove { get; }
    }
}
