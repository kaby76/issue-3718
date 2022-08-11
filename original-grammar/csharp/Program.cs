// Template generated code from trgen 0.16.6

using Antlr4.Runtime;
using Antlr4.Runtime.Tree;
using System;
using System.IO;
using System.Linq;
using System.Text;
using System.Runtime.CompilerServices;

public class Program
{
    public static SimpleBooleanParser Parser { get; set; }
    public static Lexer Lexer { get; set; }
    public static ITokenStream TokenStream { get; set; }
    public static IParseTree Tree { get; set; }
    public static string StartSymbol { get; set; } = "parse";
    public static string Input { get; set; }
    public static void SetupParse2(string input)
    {
        ICharStream str = new AntlrInputStream(input);
                var lexer = new SimpleBooleanLexer(str);
        Lexer = lexer;
        var tokens = new CommonTokenStream(lexer);
        TokenStream = tokens;
        var parser = new SimpleBooleanParser(tokens);
        Parser = parser;
        var listener_lexer = new ErrorListener<int>();
        var listener_parser = new ErrorListener<IToken>();
        lexer.RemoveErrorListeners();
        parser.RemoveErrorListeners();
        lexer.AddErrorListener(listener_lexer);
		parser.AddErrorListener(listener_parser);
    }
    public static IParseTree Parse2()
    {
        var tree = Parser.parse();
        Input = Lexer.InputStream.ToString();
        TokenStream = Parser.TokenStream;
        Tree = tree;
        return tree;
    }
    public static IParseTree Parse(string input)
    {
        ICharStream str = new AntlrInputStream(input);
            var lexer = new SimpleBooleanLexer (str);
        Lexer = lexer;
        var tokens = new CommonTokenStream(lexer);
        TokenStream = tokens;
        var parser = new SimpleBooleanParser (tokens);
        Parser = parser;
        var listener_lexer = new ErrorListener< int > ();
        var listener_parser = new ErrorListener< IToken > ();
        lexer.RemoveErrorListeners();
        parser.RemoveErrorListeners();
        lexer.AddErrorListener(listener_lexer);
        parser.AddErrorListener(listener_parser);
        var tree = parser.parse ();
        Input = lexer.InputStream.ToString();
        TokenStream = parser.TokenStream;
        Tree = tree;
        return tree;
    }

    static void Main(string[] args)
    {
        bool show_profile = false;
        bool show_tree = false;
        bool show_tokens = false;
        bool old = false;
        bool two_byte = false;
        string file_name = null;
        string input = null;
        System.Text.Encoding encoding = null;
        for (int i = 0; i < args.Length; ++i)
        {
            if (args[i].Equals("-profile"))
            {
                show_profile = true;
                continue;
            }
            else if (args[i].Equals("-tokens"))
            {
                show_tokens = true;
                continue;
            }
            else if (args[i].Equals("-two-byte"))
            {
                two_byte = true;
                continue;
            }
            else if (args[i].Equals("-old"))
            {
                old = true;
                continue;
            }
            else if (args[i].Equals("-tree"))
            {
                show_tree = true;
                continue;
            }
            else if (args[i].Equals("-input"))
                input = args[++i];
            else if (args[i].Equals("-file"))
                file_name = args[++i];
            else if (args[i].Equals("-encoding"))
            {
                ++i;
                encoding = Encoding.GetEncoding(
                    args[i],
                    new EncoderReplacementFallback("(unknown)"),
                    new DecoderReplacementFallback("(error)"));
                if (encoding == null)
                    throw new Exception(@"Unknown encoding. Must be an Internet Assigned Numbers Authority (IANA) code page name. https://www.iana.org/assignments/character-sets/character-sets.xhtml");
            }
        }
        ICharStream str = null;
        if (input == null && file_name == null)
        {
            str = CharStreams.fromStream(System.Console.OpenStandardInput());
        } else if (input != null)
        {
            str = CharStreams.fromString(input);
        } else if (file_name != null)
        {
            if (two_byte)
                str = new TwoByteCharStream(file_name);
            else if (old)
            {
                FileStream fs = new FileStream(file_name, FileMode.Open);
                str = new Antlr4.Runtime.AntlrInputStream(fs);
            }
            else if (encoding == null)
                str = CharStreams.fromPath(file_name);
            else
                str = CharStreams.fromPath(file_name, encoding);
        }
        var lexer = new SimpleBooleanLexer(str);
        if (show_tokens)
        {
            StringBuilder new_s = new StringBuilder();
            for (int i = 0; ; ++i)
            {
                var ro_token = lexer.NextToken();
                var token = (CommonToken)ro_token;
                token.TokenIndex = i;
                new_s.AppendLine(token.ToString());
                if (token.Type == Antlr4.Runtime.TokenConstants.EOF)
                    break;
            }
            System.Console.Error.WriteLine(new_s.ToString());
            lexer.Reset();
        }
        var tokens = new CommonTokenStream(lexer);
        var parser = new SimpleBooleanParser(tokens);
        var listener_lexer = new ErrorListener<int>();
        var listener_parser = new ErrorListener<IToken>();
        lexer.RemoveErrorListeners();
        parser.RemoveErrorListeners();
        lexer.AddErrorListener(listener_lexer);
        parser.AddErrorListener(listener_parser);
        DateTime before = DateTime.Now;
        if (show_profile)
        {
                parser.Profile = true;
        }
        var tree = parser.parse();
        DateTime after = DateTime.Now;
        System.Console.Error.WriteLine("Time: " + (after - before));
        if (listener_lexer.had_error || listener_parser.had_error)
        {
            System.Console.Error.WriteLine("Parse failed.");
        }
        else
        {
            System.Console.Error.WriteLine("Parse succeeded.");
        }
        if (show_tree)
        {
            System.Console.Error.WriteLine(tree.ToStringTree(parser));
        }
        if (show_profile)
        {
                System.Console.Out.WriteLine(String.Join(",\n\r", parser.ParseInfo.getDecisionInfo().Select(d => d.ToString())));
        }
        System.Environment.Exit(listener_lexer.had_error || listener_parser.had_error ? 1 : 0);
    }
}
