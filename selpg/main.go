package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type selpg_args struct {
	start_page  int
	end_page    int
	in_filename string
	page_len    int    /* default value, can be overriden by "-l number" on command line */
	page_type   string /* 'l' for lines-delimited, 'f' for form-feed-delimited */
	/* default is 'l' */
	print_dest string
}
type sp_args selpg_args

const MaxUint = ^uint(0)
const intmaxplus1 = int(MaxUint>>1) - 1

var progname = "selpg" /* program name, for error messages */

/*================================= main()=== =====================*/
func main() {
	av := os.Args
	ac := len(av)

	var sa selpg_args
	sa = selpg_args{start_page: -1, end_page: -1, in_filename: "", page_len: 72, page_type: "l", print_dest: ""}

	progname = av[0]
	_ = progname

	process_args(ac, av, &sa)
	process_input(sa)

}

/*================================= usage() =======================*/
func usage() {
	fmt.Printf("\nUSAGE: %s -sstart_page -eend_page [ -f | -llines_per_page ] [ -ddest ] [ in_filename ]\n", progname)
}

/*================================= process_args() ================*/
func process_args(ac int, av []string, psa *selpg_args) {
	var s1 string /* temp str */
	var s2 string /* temp str */
	var argno int /* arg # currently being processed */
	/* arg at index 0 is the command name itself (selpg),
	   first actual arg is at index 1,
	   last arg is at index (ac - 1) */
	var i int

	var message string
	message = progname + ": "

	/* check the command-line arguments for validity */
	if ac < 3 { /* Not enough args, minimum command is "selpg -sstartpage -eend_page"  */
		message = message + "not enough arguments\n"
		panic(message)
		usage()
	}

	/* handle mandatory args first */
	s1 = av[1] /* !!! PBO */
	if !strings.HasPrefix(s1, "-s") {
		message = message + "1st arg should be -sstart_page\n"
		panic(message)
		usage()
	}
	i, err := strconv.Atoi(s1[2:len(s1)])
	_ = err
	if err != nil || i < 1 || i > intmaxplus1 {
		message = message + "invalid start page " + s1[2:len(s1)] + "\n"
		panic(message)
		usage()
	}
	(*psa).start_page = i

	/* handle 2nd arg - end page */
	s1 = av[2] /* !!! PBO */
	if !strings.HasPrefix(s1, "-e") {
		message = message + "2nd arg should be -eend_page\n"
		panic(message)
		usage()
	}
	i, err = strconv.Atoi(s1[2:len(s1)])
	_ = err
	if err != nil || i < 1 || i > intmaxplus1 || i < (*psa).start_page {
		usage()
		message = message + "invalid end page " + s1[2:len(s1)] + "\n"
		panic(message)
	}
	(*psa).end_page = i

	/* now handle optional args */
	argno = 3
	for argno <= (ac-1) && av[argno][0] == '-' {
		s1 = av[argno]
		switch s1[1] {
		case 'l':
			s2 = s1[2:len(s1)]
			i, err = strconv.Atoi(s2)
			if err != nil || i < 1 || i > intmaxplus1 {
				message = message + "invalid page length " + s2 + "\n"
				panic(message)
				usage()
			}
			(*psa).page_len = i
			argno = argno + 1
			break

		case 'f':
			if !strings.HasPrefix(s1, "-f") {
				message = message + "option should be \"-f\"\n"
				panic(message)
				usage()
			}
			(*psa).page_type = "f"
			argno = argno + 1
			break

		case 'd':
			s2 = s1[2:len(s1)]
			if len(s2) < 1 {
				message = message + "-d option requires a printer destination\n"
				panic(message)
				usage()
			}
			(*psa).print_dest = s2
			argno = argno + 1
			break

		default:
			usage()
			message = message + "unknown option " + s1 + "\n"
			panic(message)
		} /* end switch */
	} /* end while */

	/*++argno;*/
	if argno <= ac-1 { /* there is one more arg */
		(*psa).in_filename = av[argno] /* !!! PBO */
		/* check if file exists */
		_, ero := os.Open(av[argno])
		if ero != nil {
			usage()
			message = message + "input file \"" + av[argno] + "\" does not exist\n"
			panic(message)
		}
	}
}

/*================================= process_input() ===============*/
func process_input(sa selpg_args) {
	var fioRea *bufio.Reader
	var foutWir *bufio.Writer
	var file *os.File
	var line_ctr int /* line counter */
	var page_ctr int /* page counter */
	var err error
	var message string

	/* set the input source */
	if sa.in_filename == "" {
		// fmt.Printf("START\n")
		fioRea = bufio.NewReader(os.Stdin)
	} else {
		file, err = os.Open(sa.in_filename)
		if err != nil {
			message = message + "could not open input file \"" + sa.in_filename + "\"\n"
			panic(message)
			usage()
		}
		fioRea = bufio.NewReader(file)
	}

	another := exec.Command("cat", "-n")
	in, _ := another.StdinPipe()
	if sa.print_dest != "" {
		file, err = os.OpenFile(sa.print_dest, os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			//
		}
		foutWir = bufio.NewWriter(file)

	} else {
		foutWir = bufio.NewWriter(os.Stdout)
	}

	/* begin one of two main loops based on page type */
	if sa.page_type == "l" {
		line_ctr = 0
		page_ctr = 1

		for true {
			crc, err := fioRea.ReadString('\n')
			if io.EOF == err || err != nil {
				break
			}
			line_ctr = line_ctr + 1
			if line_ctr > sa.page_len {
				page_ctr = page_ctr + 1
				line_ctr = 1
			}
			if (page_ctr >= sa.start_page) && (page_ctr <= sa.end_page) {
				if sa.print_dest != "" {
					in.Write([]byte(crc))
				} else {
					foutWir.WriteString(crc)
					foutWir.Flush()
				}
			}
		}
		if sa.print_dest != "" {
			in.Close()
			another.Stdout = os.Stdout
			another.Start()
		}
	} else {
		page_ctr = 1
		for true {
			c, _, err := fioRea.ReadRune()
			if err != nil || err == io.EOF { /* error or EOF */
				break
			}
			if c == '\f' { /* form feed */
				page_ctr = page_ctr + 1
			}
			if (page_ctr >= sa.start_page) && (page_ctr <= sa.end_page) {
				fmt.Printf("%c", c)
			}
		}
		fmt.Printf("\n")
	}
	/* end main loop */
	if page_ctr < sa.start_page {
		//message = message + ": start_page\"" + sa.start_page + "\"greater than total pages\"" + page_ctr + "\"\n"
		//panic(message)
	} else if page_ctr < sa.end_page {
		//message = message + ": end_page\"" + sa.end_page + "\"greater than total pages\"" + page_ctr + "\"\n"
		//panic(message)
	}
	return
}

/*================================= EOF ===========================*/
