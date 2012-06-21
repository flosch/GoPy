#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Quelle:
# https://de.wikibooks.org/wiki/Primzahlen:_Programmbeispiele#Python_2

def main():
    obergrenze = 10000
 
    # Jede Zahl zwischen 1 und obergrenze wird zuerst als prim angenommen
    zahlen = [True]*(obergrenze+1)
    # Da 0 und 1 keine Primzahlen sind, werden sie auf False gesetzt
    zahlen[0] = False
    zahlen[1] = False
 
    i = 2
 
    while i*i <= obergrenze:
        if zahlen[i] == True: # Die Zahl i ist als prim markiert
            # Streiche alle Vielfachen von i
            for k in range(i*i,obergrenze+1,i):
                zahlen[k] = False
 
        i = i+1
 
    # Ausgabe aller gefundenen Zahlen
    for i, v in enumerate(zahlen):
        if v == True:
            print i, 'ist prim.'
 
    return 0
 
main()
