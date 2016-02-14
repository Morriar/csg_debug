package sim.reactor;

import java.util.List;
import sim.reactor.items.Coolant;
import sim.reactor.items.ReactorRod;

public class Reactor {

    public Integer heatCapacity;
    public ReactorGrid grid;

    public Reactor(Integer heatCapacity, Integer cols, Integer rows) {
        this.heatCapacity = heatCapacity;
        this.grid = new ReactorGrid(cols, rows);
    }

    public Integer energyOutput() {
        Integer energy = 0;
        for (ReactorSlot slot : grid.slots) {
            if (slot.isEmpty()) {
                continue;
            }
            ReactorItem item = slot.getItem();
            if (item instanceof ReactorRod) {
                List<ReactorSlot> neighbors = grid.getNeighborhood(slot);
                energy += ((ReactorRod) item).clusterEnergy(neighbors);
            }
        }
        return energy;
    }

    public Integer heatOutput() {
        Integer heat = 0;
        for (ReactorSlot slot : grid.slots) {
            if (slot.isEmpty()) {
                continue;
            }
            ReactorItem item = slot.getItem();
            if (item instanceof ReactorRod) {
                List<ReactorSlot> neighbors = grid.getNeighborhood(slot);
                heat += ((ReactorRod) item).clusterHeat(neighbors);
            } else if (item instanceof Coolant) {
                heat -= ((Coolant) item).heatAbsorption();
            }
        }
        return heat + 10;
    }
}