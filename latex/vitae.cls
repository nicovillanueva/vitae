%% Custom class, based off ModernCV, by Xavier Danaux
% https://www.ctan.org/tex-archive/macros/latex/contrib/moderncv/
% Requires package: fontawesome

\NeedsTeXFormat{LaTeX2e}
\ProvidesClass{vitae}[2019/05/15]
\LoadClass{article}

\DeclareOption{a4paper}{
    \setlength\paperheight{297mm}
    \setlength\paperwidth{210mm}
}

\ExecuteOptions{a4paper,11pt,final}
\ProcessOptions\relax

\RequirePackage[utf8]{inputenc}
\RequirePackage[margin=1.5cm]{geometry}
\RequirePackage{array, xcolor}


\RequirePackage[utf8]{inputenc}
\RequirePackage{titlesec}
\RequirePackage{fontawesome}
\RequirePackage{etoolbox}
\RequirePackage{url}
\urlstyle{tt}


\definecolor{lightgray}{gray}{0.8}
\newcolumntype{L}{>{\raggedleft}p{0.2\textwidth}}
\newcolumntype{R}{p{0.8\textwidth}}
\newcommand{\VRule}{\color{lightgray}\vrule width 0.6pt}

\linespread{1.2}
% \renewcommand{\sfdefault}{pag}
\renewcommand{\familydefault}{\sfdefault}

\AtEndPreamble{
    \pagenumbering{arabic}
    \RequirePackage[unicode]{hyperref}
    \hypersetup{
        pdfborder     = 0 0 0,
        pdfauthor     = {Nicolas Villanueva},
        pdftitle      = {Nico's resum\'{e}},
        pdfkeywords   = {curriculum vit\ae{}, resum\'{e}}
    }
}

\titleformat{\section}
    {\Large\scshape\centering\ttfamily}
    {}{0em}
    {}
    [\titlerule]

\newcommand{\makeheader}{
    \maketitle
    \begin{minipage}{0.48\textwidth}
        \centering{\small\faLinkedin}~\href{https://www.linkedin.com/in/villanuevanicolas}{villanuevanicolas}\\
        {\small\faGithub}~\href{https://github.com/nicovillanueva}{nicovillanueva}
    \end{minipage}
    {\VRule}
    \begin{minipage}{0.48\textwidth}
        \centering{\small\faEnvelopeO}~nicovillanueva@pm.me\\
        {\small\faPhone}~+44 7946 479742
    \end{minipage}
}

\renewcommand*{\maketitle}{%
    \null
    \vskip 2em%
    \begin{center}%
        {\bfseries\Huge Nicolas S. Villanueva}\\
        {\vspace{1em}}
        {\fontshape{sl}\Large \{Sysadmin~{\vline}~dev\} for fun and profit}
    \end{center}%
    \vskip 1.5em
}

% \newenvironment{keyskills}{
%     \newcolumntype{L}{>{\raggedleft}p{0.2\textwidth}}
%     \newcolumntype{R}{p{0.8\textwidth}}
% }