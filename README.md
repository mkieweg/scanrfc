# scanrfc - Generate .bib files for IETF RFCs

While writing my thesis I needed something to quickly fetch bibliography information for IETF RFCs. Since I didn't find a solution that worked for me, I built one. Here it is.

```sh
$ rfc scan example

$ head rfc.bib
@misc{rfc6020,
        series =        {Request for Comments},
        number =        6020,
        howpublished =  {RFC 6020},
[...]
% Datatracker information for RFCs on the Legacy Stream is unfortunately often
% incorrect. Please correct the bibtex below based on the information in the
% actual RFC at https://rfc-editor.org/rfc/rfc1337.txt

@misc{rfc1337,
        series =        {Request for Comments},
        number =        1337,
        howpublished =  {RFC 1337},
        publisher =     {RFC Editor},
[...]
```

## Is it any good?

[Yes.](https://news.ycombinator.com/item?id=3067434)

## Usage

scanrfc can either `scan` a directory for citation keys and generate a .bib file or `fetch` RFCs that are passed as arguments.

### `rfc scan`

To scan text files (.md or .tex) for citation keys citing an RFC either as `[@rfc8040]` or `\cite{rfc7430}` and writes the bibliography information found for each RFC to `rfc.bib`. If requests the BibTex information provided by the IETF datatracker (for example `https://datatracker.ietf.org/doc/rfc7430/bibtex/`), if that's not present it fetches the respective JSON information from `https://datatracker.ietf.org/doc/$RFC/doc.json`. If no information for an RFC is found an error is logged, `rfc.bib` is created for all entries that were found.

### `rfc fetch`

Bibliography information can be fetched and logged to stdout using the `fetch` command:

```
$ rfc fetch rfc1337
% Datatracker information for RFCs on the Legacy Stream is unfortunately often
% incorrect. Please correct the bibtex below based on the information in the
% actual RFC at https://rfc-editor.org/rfc/rfc1337.txt

@misc{rfc1337,
        series =        {Request for Comments},
        number =        1337,
        howpublished =  {RFC 1337},
        publisher =     {RFC Editor},
        doi =           {10.17487/RFC1337},
        url =           {https://rfc-editor.org/rfc/rfc1337.txt},
        author =        {Robert T. Braden},
        title =         {{TIME-WAIT Assassination Hazards in TCP}},
        pagetotal =     11,
        year =          1992,
        month =         may,
        abstract =      {This note describes some theoretically-possible failure modes for TCP connections and discusses possible remedies. In particular, one very simple fix is identified. This memo provides information for the Internet community. It does not specify an Internet standard.},
}
```

## Configutaion

The output file is saved as `rfc.bib` to the current working directory. If a `.scanrfc.yaml` file is found in either the user's home directory or in the working directory is is used by `scanrfc`. The output directory can be configured using this file and the `bib-file` key.
