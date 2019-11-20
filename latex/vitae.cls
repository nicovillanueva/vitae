\NeedsTeXFormat{LaTeX2e}
\ProvidesClass{vitae}[2019/05/15]
\LoadClass{article}

\DeclareOption{a4paper}{
    \setlength\paperheight{297mm}
    \setlength\paperwidth{210mm}
}

\ExecuteOptions{a4paper}
\ProcessOptions\relax

\RequirePackage[utf8]{inputenc}
\RequirePackage[margin=1.5cm]{geometry}
\RequirePackage{array, xcolor}
\RequirePackage{tikz}
\RequirePackage{enumitem}
\RequirePackage{titlesec}
\RequirePackage{fontawesome}
\RequirePackage{etoolbox}
\RequirePackage{url}
\RequirePackage{multicol}

\urlstyle{tt}

\RequirePackage{monokai}
\RequirePackage{vitae-style}

\AtEndPreamble{
    \RequirePackage[unicode]{hyperref}
    \hypersetup{
        urlcolor      = \@urlColor,
        colorlinks    = true,
        pdfborder     = 0 0 0,
        pdfauthor     = {Nicolas Villanueva},
        pdftitle      = {Nico's resum\'{e}},
        pdfkeywords   = {curriculum vit\ae{}, resum\'{e}}
    }
}

%% Draw a progress bar. <param>/10
%% Usage: \progressbar{3} => Renders a bar 30% filled
\newcommand{\progressbar}[1]{
    \begin{tikzpicture}[xscale=0.1]
        \filldraw[color=\@progressBackColor](0,0) rectangle (10,0.2);
        \filldraw[color=\@progressFillColor] (0,0) rectangle (#1,0.2);
    \end{tikzpicture}
}

%% Tri ▲
\newcommand{\tri}{
    \begin{tikzpicture}[scale=.5]
        \fill[color=\@triColor] (0,0) +(0:1) -- +(60:1) -- +(60:0) -- cycle;
    \end{tikzpicture}
}

%% Revamp the title command, used in \makeheader
\renewcommand*{\maketitle}{%
    \null
    \vskip 2em
    \begin{center}%
        {\bfseries\scshape\Huge {\@titleNameStyle{Nicolas}}{\@titleLastStyle{Villanueva}}}
        \vskip 1em
        {\@titleSubStyle{\{Sysadmin~{\vline}~dev\} for fun and profit}}
        % \vskip 1.5em 
        % \tri
    \end{center}%
    \vskip 1.5em
}

%% Insert the header, with contact info
\newcommand{\makeheader}{
    \maketitle
    \begin{minipage}{0.49\textwidth}
        \centering{\small\faLinkedin}~\href{https://www.linkedin.com/in/villanuevanicolas}{\@code{villanuevanicolas}}

        {\small\faGithub}~\href{https://github.com/nicovillanueva}{\@code{nicovillanueva}}
    \end{minipage}
    \hskip 2pt  % TODO: Fix this mf
    {\VRule}
    \begin{minipage}{0.49\textwidth}
        \centering{\small\faEnvelopeO}~nicovillanueva@pm.me

        {\small\faPhone}~+44 7946 479742 % (UK); +549 11 3151 8837 (AR)

        {\small\faMapMarker}~SE12, London, UK
    \end{minipage}
}
